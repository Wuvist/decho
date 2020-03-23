//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Wuvist/decho/controller"
)

func getWebApp() (*WebApp, error) {
	wire.Build(newEcho, wire.Struct(new(WebApp), "*"), loadTomlConf, newMsg,
		controller.NewBlogController,
		controller.NewCateController,
		controller.NewHomeController)
	return nil, nil
}
