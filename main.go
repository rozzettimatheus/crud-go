package main

import (
	"log"

	"github.com/gin-gonic/gin"
	env "github.com/joho/godotenv"
	"github.com/rozzettimatheus/crud-go/src/controller"
	"github.com/rozzettimatheus/crud-go/src/controller/routes"
	"github.com/rozzettimatheus/crud-go/src/model/service"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
