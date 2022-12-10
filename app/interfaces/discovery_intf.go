package interfaces

import (
	"monica/errcode"
	"net/http"
)

// Fetch 获取服务实例
func (s *MonicaHttpServiceImpl) Fetch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ns := query.Get("ns")
	sname := query.Get("sname")
	nodeList, err := s.disv.Fetch(r.Context(), ns, sname)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, nodeList)
}

// Poll 获取服务实例（长轮询）
func (s *MonicaHttpServiceImpl) Poll(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ns := query.Get("ns")
	sname := query.Get("sname")
	nodeList, err := s.disv.Poll(r.Context(), ns, sname)
	if err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, nodeList)
}
