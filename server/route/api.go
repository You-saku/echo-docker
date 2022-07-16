package route

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	"app/server/controllers" // controllerを読み込む
)

func Routing() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	})

	e.GET("/sample", sample)

	// prefixつきrouting
	api := e.Group("/api")
	api.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, "api")
	})

	api.POST("/user", controllers.CreateUser)

	e.Logger.Fatal(e.Start(":80"))
}

// ルーティングに使う関数
func sample(c echo.Context) error {
	return c.String(http.StatusOK, "Sample.")
}
