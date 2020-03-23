package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CateController handles blog category requests
type CateController struct {
}

func (controller *CateController) show(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

// NewCateController return CateController bind with echo engine
func NewCateController(e *echo.Echo) (*CateController, error) {
	cate := &CateController{}
	e.GET("/cate.go", cate.show)

	return cate, nil
}
