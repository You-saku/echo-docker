package route

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// controllerを読み込む
	"app/server/controllers"
	// バリデーション
	"app/server/validates"

	// ミドルウェア
	"github.com/labstack/echo/v4/middleware"
	"app/server/token"
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

	// バリデーションを設定
	e.Validator = validates.SetValidator()

	// prefixつきrouting
	api := e.Group("/api")
	api.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, "api")
	})

	// ログイン(jwtを返す)
	api.POST("/login", controllers.Login)

	// ログインしてたら叩けるエンドポイントを作る
	login := api

	// 認証設定
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(token.Key),
	}
	// メモ
	// このJWTWithConfigを使うと間違ったJWTなら401, Authorization headerがなければ400になる
	login.Use(middleware.JWTWithConfig(config))
	// userのCRUD
	login.POST("/user", controllers.CreateUser)
	login.GET("/users", controllers.GetUsers)
	login.GET("/users/:id", controllers.GetUser)
	login.PUT("/users/:id", controllers.UpdateUser)
	login.DELETE("/users/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":80"))
}

// ルーティングに使う関数
func sample(c echo.Context) error {
	return c.String(http.StatusOK, "Sample.")
}
