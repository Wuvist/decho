package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	InitEngine(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func newMsg() string {
	return "bingo"
}
