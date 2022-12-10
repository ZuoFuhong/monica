package web

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"monica/app/interfaces"
	"net/http"
	"net/url"
	"strings"
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

func (r *Router) registerHandler(s *interfaces.MonicaHttpServiceImpl) {
	r.Handle("/ping", r.ch.ThenFunc(s.Ping)).Methods("GET")
	r.Handle("/api/v1/services", r.ch.ThenFunc(s.DoServiceList)).Methods("GET")
	r.Handle("/api/v1/service", r.ch.ThenFunc(s.DoSaveService)).Methods("POST")
	r.Handle("/api/v1/service", r.ch.ThenFunc(s.DoDeleteService)).Methods("DELETE")
	r.Handle("/api/v1/service_namespaces", r.ch.ThenFunc(s.DoServiceNamespaceList)).Methods("GET")
	r.Handle("/api/v1/service_instances", r.ch.ThenFunc(s.DoServiceInstanceList)).Methods("GET")
	r.Handle("/api/v1/service_instance", r.ch.ThenFunc(s.DoSaveServiceInstance)).Methods("POST")
	r.Handle("/api/v1/service_instance", r.ch.ThenFunc(s.DoDeleteServiceInstance)).Methods("DELETE")
	// 服务注册/发现
	r.Handle("/api/v1/register", r.ch.ThenFunc(s.Register)).Methods("POST")
	r.Handle("/api/v1/renew", r.ch.ThenFunc(s.Renew)).Methods("POST")
	r.Handle("/api/v1/cancel", r.ch.ThenFunc(s.Cancel)).Methods("POST")
	r.Handle("/api/v1/fetch", r.ch.ThenFunc(s.Fetch)).Methods("GET")
	r.Handle("/api/v1/poll", r.ch.ThenFunc(s.Poll)).Methods("GET")
	// 静态资源
	r.PathPrefix("/").Handler(vueRouter("/", http.FileServer(http.Dir("./dist"))))
}

func vueRouter(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := strings.TrimPrefix(r.URL.Path, prefix)
		// 适配 Vue 路由
		if strings.HasPrefix(p, "service") || strings.HasPrefix(p, "naming-service") || strings.HasPrefix(p, "kvgroup") {
			p = ""
		}
		rp := strings.TrimPrefix(r.URL.RawPath, prefix)
		if len(p) < len(r.URL.Path) && (r.URL.RawPath == "" || len(rp) < len(r.URL.RawPath)) {
			r2 := new(http.Request)
			*r2 = *r
			r2.URL = new(url.URL)
			*r2.URL = *r.URL
			r2.URL.Path = p
			r2.URL.RawPath = rp
			h.ServeHTTP(w, r2)
		} else {
			http.NotFound(w, r)
		}
	})
}
