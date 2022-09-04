package main

import (
	"log"
	"monica-admin/app/interfaces"
	"monica-admin/pkg/database"
	"monica-admin/web"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
	s := web.NewServer()
	gormDB := database.NewGormDB()
	service := interfaces.InitializeService(gormDB)
	s.Register(service)
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
