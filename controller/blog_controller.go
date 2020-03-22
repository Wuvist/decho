package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BlogController defines controller for blog module
type BlogController struct {
	msg string
}

func (controller *BlogController) hello(c echo.Context) error {
	return c.String(http.StatusOK, controller.msg)
}

// NewBlogController return BlogController bind with echo engine
func NewBlogController(e *echo.Echo, msg string) (*BlogController, error) {
	blog := &BlogController{
		msg: msg,
	}
	e.GET("/blog/hello", blog.hello)

	return blog, nil
}
