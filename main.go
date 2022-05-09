package main

import (
	"CodingTest/configs"
	"CodingTest/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"

	_adminHandler "CodingTest/delivery/handler/admin"
	_adminRepo "CodingTest/repository/admin"
	_adminUsecase "CodingTest/usecase/admin"

	"CodingTest/delivery/routes"
)

func main() {
	config := configs.GetConfig()
	db, _ := utils.Connect(config)

	adminRepo := _adminRepo.NewAdminRepository(db)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo)
	adminHandler := _adminHandler.NewAdminHandler(adminUsecase)

	e := echo.New()
	routes.AdminPath(e, adminHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
