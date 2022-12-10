package interfaces

import (
	"encoding/json"
	"monica/app/domain/entity"
	"monica/consts"
	"monica/errcode"
	"monica/pkg/log"
	"net/http"
	"time"
)

// DoServiceInstanceList 管理后台-服务实例列表
func (s *MonicaHttpServiceImpl) DoServiceInstanceList(w http.ResponseWriter, r *http.Request) {
	// 1.解析参数
	query := r.URL.Query().Get("query")
	cond, err := parseSearchInsReq(query)
	if err != nil {
		log.ErrorContextf(r.Context(), "call parseSearchInsReq failed, query = %s, err: %v", query, err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 2.分页查询
	sList, total, err := s.insv.ListInstanceByPage(r.Context(), cond.Namespace, cond.ServiceName, cond.IP, cond.Port, cond.Page, cond.PageSize)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	voList := make([]*entity.InstanceVO, 0)
	now := time.Now()
	for _, ins := range sList {
		insVO := &entity.InstanceVO{
			ID:          ins.ID,
			Namespace:   ins.Namespace,
			ServiceName: ins.ServiceName,
			Healthy:     consts.Critical,
			Isolate:     ins.Isolate,
			IP:          ins.IP,
			Port:        ins.Port,
			Weight:      ins.Weight,
			Metadata:    ins.Metadata,
			Ctime:       ins.CreatedAt.Unix(),
			Mtime:       ins.UpdatedAt.Unix(),
		}
		// 计算健康状态
		if ins.RenewedAt != nil && ins.RenewedAt.Unix()+consts.SyncNSInterval > now.Unix() {
			insVO.Healthy = consts.Passing
		}
		voList = append(voList, insVO)
	}
	data := make(map[string]interface{})
	data["list"] = voList
	data["total"] = total
	Ok(w, data)
}

// DoSaveServiceInstance 管理后台-创建/修改 服务实例
func (s *MonicaHttpServiceImpl) DoSaveServiceInstance(w http.ResponseWriter, r *http.Request) {
	req := new(SaveInstanceReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.ErrorContextf(r.Context(), "call json.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 1.检查参数
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "参数错误")
		return
	}
	// 2.保存服务实例
	if err := s.insv.SaveInstance(r.Context(), req.convertToEntity()); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, "ok")
}

// DoDeleteServiceInstance 管理后台-删除服务实例
func (s *MonicaHttpServiceImpl) DoDeleteServiceInstance(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ns := query.Get("ns")
	ip := query.Get("ip")
	sname := query.Get("sname")
	if ns == "" || ip == "" {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	if err := s.insv.DeleteInstance(r.Context(), ns, sname, ip); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, "ok")
}
