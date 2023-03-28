package controller

import (
	"net/http"

	"echo_gorm_rest_api/model"

	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	posts := []model.Post{}
	model.DB.Find(&posts)
	return c.JSON(http.StatusOK, posts)
}

func GetPost(c echo.Context) error {
	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return err
	}
	model.DB.Take(&post)
	return c.JSON(http.StatusOK, post)
}

func CreatePost(c echo.Context) error {
	post := model.Post{}
	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	model.DB.Create(&post)
	return c.JSON(http.StatusCreated, post)
}
