package main

import (
	"log"
	"monica/app/interfaces"
	"monica/pkg/database"
	"monica/web"
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
