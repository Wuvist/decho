package controller

import (
	"net/http"

	"github.com/Wuvist/decho/static"
	"github.com/labstack/echo/v4"
)

// StaticController handles static requests
type StaticController struct{}

// NewStaticController return StaticController bind with echo engine
func NewStaticController(e *echo.Echo) (*StaticController, error) {
	ctrl := &StaticController{}
	e.GET("/Template/*", echo.WrapHandler(http.FileServer(static.FS(false))))

	return ctrl, nil
}
