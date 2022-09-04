package interfaces

import (
	"net/http"
)

type AdminHttpServiceImpl struct {
}

func InitializeService() *AdminHttpServiceImpl {
	return &AdminHttpServiceImpl{}
}

func (k *AdminHttpServiceImpl) Ping(w http.ResponseWriter, r *http.Request) {
	Ok(w, "ok")
}
