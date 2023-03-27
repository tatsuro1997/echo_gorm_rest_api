package main

import (
	"echo_gorm_rest_api/controller"
	"echo_gorm_rest_api/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.GET("users", controller.GetUsers)
	e.GET("user/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}