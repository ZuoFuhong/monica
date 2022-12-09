package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"monica-admin/app/interfaces"
)

type Router struct {
	*mux.Router
	ch *alice.Chain
}

func NewRouter() *Router {
	mw := NewMiddleware()
	chain := alice.New(mw.RequestMetricHandler)
	return &Router{
		Router: mux.NewRouter(),
		ch:     &chain,
	}
}

func (r *Router) registerHandler(s *interfaces.AdminHttpServiceImpl) {
	r.Handle("/ping", r.ch.ThenFunc(s.Ping)).Methods("GET")
	r.Handle("/api/v1/services", r.ch.ThenFunc(s.DoServiceList)).Methods("GET")
	r.Handle("/api/v1/service", r.ch.ThenFunc(s.DoSaveService)).Methods("POST")
	r.Handle("/api/v1/service", r.ch.ThenFunc(s.DoDeleteService)).Methods("DELETE")
	r.Handle("/api/v1/service_instances", r.ch.ThenFunc(s.DoServiceInstanceList)).Methods("GET")
	r.Handle("/api/v1/service_instance", r.ch.ThenFunc(s.DoSaveServiceInstance)).Methods("POST")
	r.Handle("/api/v1/service_instance", r.ch.ThenFunc(s.DoDeleteServiceInstance)).Methods("DELETE")
}
