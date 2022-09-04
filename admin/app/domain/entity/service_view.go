package entity

type ServiceVO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Business string `json:"business"`
	Owners   string `json:"owners"`
	Remark   string `json:"remark"`
	Ctime    int64  `json:"ctime"`
	Mtime    int64  `json:"mtime"`
}
