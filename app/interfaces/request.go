package interfaces

import (
	"encoding/base64"
	"encoding/json"
	"monica/app/domain/entity"
)

// SearchReq 搜索服务列表
type SearchReq struct {
	Kw       string `json:"kw"`        // 服务名称
	Busi     string `json:"busi"`      // 业务名称
	Page     int    `json:"page"`      // 分页
	PageSize int    `json:"page_size"` // 单页条数
}

// parseSearchReq 解析请求参数
func parseSearchReq(query string) (*SearchReq, error) {
	bytes, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		return nil, err
	}
	req := new(SearchReq)
	err = json.Unmarshal(bytes, req)
	return req, err
}

// SearchNSReq 搜索服务命名空间
type SearchNSReq struct {
	Kw       string `json:"kw"`        // 服务名称
	IP       string `json:"ip"`        // 实例IP
	Page     int    `json:"page"`      // 分页
	PageSize int    `json:"page_size"` // 单页条数
}

// parseSearchNSReq 解析请求参数
func parseSearchNSReq(query string) (*SearchNSReq, error) {
	bytes, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		return nil, err
	}
	req := new(SearchNSReq)
	err = json.Unmarshal(bytes, req)
	return req, err
}

// SaveServiceReq 保存/修改 服务
type SaveServiceReq struct {
	ID       int    `json:"id"`       // 服务ID
	Name     string `json:"name"`     // 服务名称
	Business string `json:"business"` // 业务名称
	Owners   string `json:"owners"`   // 负责人
	Remark   string `json:"remark"`   // 备注
}

func (r *SaveServiceReq) CheckRequiredParam() bool {
	if r.Name == "" || r.Business == "" || r.Owners == "" {
		return false
	}
	return true
}

func (r *SaveServiceReq) convertToEntity() *entity.Service {
	record := &entity.Service{
		ID:       r.ID,
		Name:     r.Name,
		Business: r.Business,
		Owners:   r.Owners,
		Remark:   r.Remark,
	}
	return record
}

// SearchInsReq 搜索服务实例列表
type SearchInsReq struct {
	Namespace   string `json:"ns"`        // 命名空间
	ServiceName string `json:"sname"`     // 服务名称
	IP          string `json:"ip"`        // 实例IP
	Port        int    `json:"port"`      // 实例端口
	Page        int    `json:"page"`      // 分页
	PageSize    int    `json:"page_size"` // 单页条数
}

// parseSearchInsReq 解析请求参数
func parseSearchInsReq(query string) (*SearchInsReq, error) {
	bytes, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		return nil, err
	}
	req := new(SearchInsReq)
	err = json.Unmarshal(bytes, req)
	return req, err
}

// SaveInstanceReq 保存/修改 服务实例
type SaveInstanceReq struct {
	ID          int    `json:"id"`       // 实例ID
	Namespace   string `json:"ns"`       // 命名空间
	ServiceName string `json:"sname"`    // 服务名称
	Isolate     int    `json:"isolate"`  // 隔离状态：1-不隔离 2-隔离
	IP          string `json:"ip"`       // 实例IP
	Port        int    `json:"port"`     // 端口
	Weight      int    `json:"weight"`   // 权重
	Metadata    string `json:"metadata"` // 元数据
}

func (r *SaveInstanceReq) CheckRequiredParam() bool {
	if r.Namespace == "" || r.ServiceName == "" || r.IP == "" || r.Port == 0 {
		return false
	}
	if r.Isolate != 1 && r.Isolate != 2 {
		return false
	}
	if r.Weight < 0 || r.Weight > 100 {
		return false
	}
	return true
}

func (r *SaveInstanceReq) convertToEntity() *entity.ServiceInstance {
	return &entity.ServiceInstance{
		ID:          r.ID,
		Namespace:   r.Namespace,
		ServiceName: r.ServiceName,
		Isolate:     r.Isolate,
		IP:          r.IP,
		Port:        r.Port,
		Weight:      r.Weight,
		Metadata:    r.Metadata,
	}
}

// RegisterReq 服务注册
type RegisterReq struct {
	Token       string               `json:"token"`        // Token 令牌
	Namespace   string               `json:"namespace"`    // 命名空间
	ServiceName string               `json:"service_name"` // 服务名称
	Node        *entity.InstanceNode `json:"node"`         // 实例节点
}

func (r *RegisterReq) CheckRequiredParam() bool {
	if r.Token == "" || r.Namespace == "" || r.ServiceName == "" || r.Node == nil {
		return false
	}
	if r.Node.IP == "" {
		return false
	}
	if r.Node.Weight < 0 || r.Node.Weight > 100 {
		return false
	}
	return true
}

// RenewReq 服务更新
type RenewReq struct {
	Token       string `json:"token"`        // Token 令牌
	Namespace   string `json:"namespace"`    // 命名空间
	ServiceName string `json:"service_name"` // 服务名称
	IP          string `json:"ip"`           // 实例IP
}

func (r *RenewReq) CheckRequiredParam() bool {
	if r.Token == "" || r.Namespace == "" || r.ServiceName == "" || r.IP == "" {
		return false
	}
	return true
}

// DeregisterReq 服务注销
type DeregisterReq struct {
	Token       string `json:"token"`        // Token 令牌
	Namespace   string `json:"namespace"`    // 命名空间
	ServiceName string `json:"service_name"` // 服务名称
	IP          string `json:"ip"`           // 实例IP
}

func (r *DeregisterReq) CheckRequiredParam() bool {
	if r.Token == "" || r.Namespace == "" || r.ServiceName == "" || r.IP == "" {
		return false
	}
	return true
}
