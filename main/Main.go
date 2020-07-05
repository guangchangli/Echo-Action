package main

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// 得到一个 echo.Echo 的实例
	e := echo.New()
	// 注册路由
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// 开启 HTTP Server
	e.Logger.Fatal(e.Start(":2020"))
}