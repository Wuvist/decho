package main

import (
	"github.com/Wuvist/decho/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.BindBlogController(e)
	e.Logger.Fatal(e.Start(":1323"))
}
