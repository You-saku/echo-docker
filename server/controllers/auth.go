package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	// モデルを定義してる自作パッケージ
	"app/server/models"
	// jwt周り
	"app/server/token"
)

func Login(c echo.Context) (err error) {
	jwtClaim := new(token.JwtCustomClaims)

	c.Bind(jwtClaim)
	if err := c.Validate(jwtClaim); err != nil {
		return err
	}
	t, err := token.CreateJwt(jwtClaim.Email, jwtClaim.Password)

	if err != nil {
		return err
	}

	user := models.User{}
	result := user.Verify(jwtClaim.Email, jwtClaim.Password)

	// result.Error はDBエラーがあったらnilじゃなくなる
	// Todo 正式なエラーメッセージを取得して表示する
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "This user is not exist.",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
