package interfaces

import (
	"gorm.io/gorm"
	"monica-admin/app/domain/service"
	"monica-admin/app/infra/database"
	"net/http"
)

type AdminHttpServiceImpl struct {
	serv service.IService
	insv service.IServiceInstance
}

func InitializeService(db *gorm.DB) *AdminHttpServiceImpl {
	servrepos := database.NewIServiceRepos(db)
	insRepos := database.NewIServiceInstanceRepos(db)
	serv := service.NewService(servrepos)
	insv := service.NewServiceInstance(servrepos, insRepos)
	return &AdminHttpServiceImpl{
		serv: serv,
		insv: insv,
	}
}

func (s *AdminHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
