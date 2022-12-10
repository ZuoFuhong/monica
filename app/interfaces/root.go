package interfaces

import (
	"gorm.io/gorm"
	"monica/app/domain/service"
	"monica/app/infra/database"
	"net/http"
)

type MonicaHttpServiceImpl struct {
	serv service.IService
	insv service.IServiceInstance
	disv service.IDiscoveryService
	resv service.IRegisterService
}

func InitializeService(db *gorm.DB) *MonicaHttpServiceImpl {
	servrepos := database.NewIServiceRepos(db)
	insRepos := database.NewIServiceInstanceRepos(db)
	serv := service.NewService(servrepos, insRepos)
	insv := service.NewServiceInstance(servrepos, insRepos)
	resv := service.NewRegisterService(insRepos)
	disv := service.NewDiscoveryService(insRepos)
	return &MonicaHttpServiceImpl{
		serv: serv,
		insv: insv,
		disv: disv,
		resv: resv,
	}
}

func (s *MonicaHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
