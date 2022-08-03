package controllers

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// 文字列を数字に変換する
	"strconv"

	// モデルを定義してる自作パッケージ
	"app/server/models"
)

func CreateUser(c echo.Context) (err error) {
	user := new(models.User)

	c.Bind(user) // これ最強
	if err := c.Validate(user); err != nil {
		return err
	}

	user.Create()

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := models.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := models.User{}
	user.Show(id)

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := new(models.User)

	c.Bind(user)
	if err := c.Validate(user); err != nil {
		return err
	}

	user.Update(id)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := models.User{}
	user.Delete(id)

	return c.NoContent(http.StatusNoContent)
}
