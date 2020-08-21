package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//创建实例
	c := echo.New()
	//配置日志
	configLogger(c)
	//静态文件路由
	c.Static("img", "img")
	c.File("/favicon.ico", "img/favicon.ico")
	//设置中间
	setMiddleware(c)
	//注册路由
	RegisterRoutes(c)
	c.Start("2020")

}
func configLogger(e *echo.Echo) {
	// 定义日志级别
	e.Logger.SetLevel(log.INFO)
	// 记录业务日志
	echoLog, err := os.OpenFile("log/echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	// 同时输出到文件和终端
	e.Logger.SetOutput(io.MultiWriter(os.Stdout, echoLog))
}
func setMiddleware(c *echo.Context){

}
func RegisterRoutes(c *echo.Context){

}
