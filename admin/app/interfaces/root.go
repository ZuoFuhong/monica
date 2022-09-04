package interfaces

import (
	"gorm.io/gorm"
	"monica-admin/app/domain/service"
	"monica-admin/app/infra/database"
	"net/http"
)

type AdminHttpServiceImpl struct {
	serv service.IService
}

func InitializeService(db *gorm.DB) *AdminHttpServiceImpl {
	repos := database.NewIServiceRepos(db)
	serv := service.NewService(repos)
	return &AdminHttpServiceImpl{
		serv: serv,
	}
}

func (s *AdminHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
