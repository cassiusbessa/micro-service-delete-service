package main

import (
	"github.com/cassiusbessa/delete-service/handlers"
	"github.com/cassiusbessa/delete-service/logs"
	"github.com/cassiusbessa/delete-service/repositories"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	r := handlers.Router()
	repositories.Repo.Ping()
	r.DELETE("/services/:company/:id", handlers.DeleteService)
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}
