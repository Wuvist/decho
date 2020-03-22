//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

	"github.com/Wuvist/decho/controller"
)

func InitEngine(e *echo.Echo) (*controller.BlogController, error) {
	wire.Build(controller.NewBlogController, newMsg)
	return nil, nil
}
