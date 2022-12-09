package interfaces

import (
	"gorm.io/gorm"
	"monica/app/domain/service"
	"monica/app/infra/database"
	"net/http"
)

type AdminHttpServiceImpl struct {
	serv service.IService
	insv service.IServiceInstance
}

func InitializeService(db *gorm.DB) *AdminHttpServiceImpl {
	servrepos := database.NewIServiceRepos(db)
	insRepos := database.NewIServiceInstanceRepos(db)
	serv := service.NewService(servrepos, insRepos)
	insv := service.NewServiceInstance(servrepos, insRepos)
	return &AdminHttpServiceImpl{
		serv: serv,
		insv: insv,
	}
}

func (s *AdminHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
