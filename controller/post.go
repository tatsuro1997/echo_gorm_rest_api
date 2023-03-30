package controller

import (
	"net/http"
	"strconv"

	"github.com/tatsuro1997/echo_gorm_rest_api/model"

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

	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	post.UserID = uint(userID)

	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := model.DB.Create(&post)
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error.Error())
	}
	return c.JSON(http.StatusCreated, post)
}

func UpdatePost(c echo.Context) error {
	post := model.Post{}

	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	model.DB.Where("ID = ?", postID).First(&post)

	if err := c.Bind(&post); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res := model.DB.Model(&post).Where("ID", postID).Updates(&post)
	if res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error.Error())
	}

	return c.JSON(http.StatusOK, post)
}
