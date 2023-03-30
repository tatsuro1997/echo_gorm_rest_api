package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tatsuro1997/echo_gorm_rest_api/controller"
	"github.com/tatsuro1997/echo_gorm_rest_api/model"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.GET("users", controller.GetUsers)
	e.GET("users/:id", controller.GetUser)
	e.POST("users", controller.CreateUser)
	e.PATCH("users/:id", controller.UpdateUser)
	e.DELETE("users/:id", controller.DeleteUser)

	e.GET("posts", controller.GetPosts)
	e.GET("posts/:id", controller.GetPost)
	e.POST("users/:user_id/posts", controller.CreatePost)
	e.PATCH("users/:user_id/posts/:id", controller.UpdatePost)
	e.Logger.Fatal(e.Start(":8080"))
}
