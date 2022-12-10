package interfaces

import (
	"encoding/json"
	"monica/errcode"
	"monica/pkg/log"
	"net/http"
)

// Register 服务注册
func (s *MonicaHttpServiceImpl) Register(w http.ResponseWriter, r *http.Request) {
	req := new(RegisterReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.ErrorContextf(r.Context(), "call json.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 检查参数
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	// 令牌检查
	if err := s.serv.VerifyServiceToken(r.Context(), req.Token, req.Namespace, req.ServiceName); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	// 服务注册
	if err := s.resv.Register(r.Context(), req.Namespace, req.ServiceName, req.Node); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, "ok")
}

// Renew 服务刷新
func (s *MonicaHttpServiceImpl) Renew(w http.ResponseWriter, r *http.Request) {
	req := new(RenewReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.ErrorContextf(r.Context(), "call json.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 检查参数
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	// 令牌检查
	if err := s.serv.VerifyServiceToken(r.Context(), req.Token, req.Namespace, req.ServiceName); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	if err := s.resv.Renew(r.Context(), req.Namespace, req.ServiceName, req.IP); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	Ok(w, "ok")
}

// Cancel 服务注销
func (s *MonicaHttpServiceImpl) Cancel(w http.ResponseWriter, r *http.Request) {
	req := new(CancelReq)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.ErrorContextf(r.Context(), "call json.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "参数解析失败")
		return
	}
	// 检查参数
	if !req.CheckRequiredParam() {
		Error(w, errcode.BadRequestParam, "缺少参数")
		return
	}
	// 令牌检查
	if err := s.serv.VerifyServiceToken(r.Context(), req.Token, req.Namespace, req.ServiceName); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
		return
	}
	if err := s.resv.Cancel(r.Context(), req.Namespace, req.ServiceName, req.IP); err != nil {
		Error(w, errcode.InternalServerError, err.Error())
	}
	Ok(w, "ok")
}
