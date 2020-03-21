package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BlogController defines controller for blog module
type BlogController struct{}

func (controller *BlogController) hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// BindBlogController binds BlogController to echo engine
func BindBlogController(e *echo.Echo) error {
	blog := BlogController{}
	e.GET("/blog/hello", blog.hello)
	return nil
}
