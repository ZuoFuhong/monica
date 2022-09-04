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
}
