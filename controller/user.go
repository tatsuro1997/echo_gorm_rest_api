package controller

import (
	"log"
	"net/http"

	"echo_gorm_rest_api/model"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err!= nil {
		return err
	}
	model.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err!= nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func UpdateUser(c echo.Context) error {
	user := model.User{}
  if err := c.Bind(&user); err!= nil {
    return err
  }
  res := model.DB.Save(&user)
	if res.Error!= nil {
		log.Println(res.Error)
	}
  return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
  if err := c.Bind(&user); err!= nil {
    return err
  }
  res := model.DB.Delete(&user)
  if res.Error!= nil {
    log.Println(res.Error)
  }
  return c.JSON(http.StatusOK, user)
}
