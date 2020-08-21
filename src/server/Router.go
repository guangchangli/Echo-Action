package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)
type user struct {
	name string `json:"Name"`
	age int `json:"Age"`
	tel string `json:"phone"`
}


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/getUser/:id",getUserFromPath)
	e.GET("/getUserParam",getUserFromParam)
	e.Logger.Fatal(e.Start(":2020"))
}
func getUserFromPath(c echo.Context) error {
	param := c.Param("id")
	c.Logger().Printf("%s\n",param)
	marshal, _ := json.Marshal(user{
		name: "echo",
		age:  1,
		tel:  "https://www.bookstack.cn/read/echo-v3-zh/index-index.md",
	})
	fmt.Println(marshal)
	return  c.String(http.StatusOK, param)
}
func getUserFromParam(c echo.Context)user{
	name:= c.QueryParam("name")
	tel := c.QueryParam("tel")
	log.Fatalf("getUserFromParam请求参数为：name:[%s],tel:[%s]\n",name,tel)
	return user{
		name: "echo",
		age:  1,
		tel:  "https://www.bookstack.cn/read/echo-v3-zh/index-index.md",
	}
}
