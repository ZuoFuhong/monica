package interfaces

import (
	"encoding/json"
	"monica-admin/app/domain/entity"
	"monica-admin/errcode"
	"monica-admin/pkg/log"
	"net/http"
)

// DoServiceList 管理后台-服务列表
func (s *AdminHttpServiceImpl) DoServiceList(w http.ResponseWriter, r *http.Request) {
	// 1.解析参数
	query := r.URL.Query().Get("query")
	cond, err := parseSearchReq(query)
	if err != nil {
		log.ErrorContextf(r.Context(), "call parseSearchReq failed, query = %s, err: %v", query, err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 2.分页查询
	sList, total, err := s.serv.ListServiceByPage(r.Context(), cond.Kw, cond.Busi, cond.Page, cond.PageSize)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	svoList := make([]*entity.ServiceVO, 0)
	for _, sdo := range sList {
		var mtime int64
		if sdo.UpdatedAt != nil {
			mtime = sdo.UpdatedAt.Unix()
		}
		svo := &entity.ServiceVO{
			Id:       sdo.ID,
			Name:     sdo.Name,
			Business: sdo.Business,
			Owners:   sdo.Owners,
			Remark:   sdo.Remark,
			Ctime:    sdo.CreatedAt.Unix(),
			Mtime:    mtime,
		}
		svoList = append(svoList, svo)
	}
	data := make(map[string]interface{})
	data["list"] = svoList
	data["total"] = total
	Ok(w, data)
}

// DoSaveService 管理后台-创建/修改 服务
func (s *AdminHttpServiceImpl) DoSaveService(w http.ResponseWriter, r *http.Request) {
	req := new(SaveServiceReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.ErrorContextf(r.Context(), "call json.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 1.检查参数
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "缺少必传参数")
		return
	}
	// 2.保存服务
	if err := s.serv.SaveService(r.Context(), req.convertToEntity()); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	// 3.回包
	Ok(w, "ok")
}

// DoDeleteService 管理后台-删除服务
func (s *AdminHttpServiceImpl) DoDeleteService(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		Error(w, errcode.BadRequestParam, "无效的参数")
		return
	}
	if err := s.serv.DeleteService(r.Context(), name); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, "ok")
}
