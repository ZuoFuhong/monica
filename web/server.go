package web

import (
	"fmt"
	"monica/app/interfaces"
	"monica/pkg/config"
	"monica/pkg/log"
	"net/http"
)

type Server struct {
	name   string
	addr   string
	router *Router
}

func NewServer() *Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("load config fail: " + err.Error())
	}
	config.SetGlobalConfig(cfg)

	return &Server{
		name:   cfg.Server.Name,
		addr:   fmt.Sprintf("%s:%d", cfg.Server.Addr, cfg.Server.Port),
		router: NewRouter(),
	}
}

func (s *Server) Register(service *interfaces.MonicaHttpServiceImpl) {
	s.router.registerHandler(service)
}

func (s *Server) Serve() error {
	log.Debugf("%s runs on http://%s", s.name, s.addr)
	return http.ListenAndServe(s.addr, s.router)
}
