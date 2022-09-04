package interfaces

import (
	"encoding/base64"
	"encoding/json"
	"monica-admin/app/domain/entity"
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

// SaveServiceReq 保存/修改 服务
type SaveServiceReq struct {
	ID       int64  `json:"id"`       // 服务ID
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
		Id:       r.ID,
		Name:     r.Name,
		Business: r.Business,
		Owners:   r.Owners,
		Remark:   r.Remark,
	}
	return record
}
