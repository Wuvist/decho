//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Wuvist/decho/controller"
)

var dbProviders = wire.NewSet(
	getBloggerDB,
	getArticalDB,
	getCategoryDB,
)

func getWebApp() (*WebApp, error) {
	wire.Build(newEcho, wire.Struct(new(WebApp), "*"), loadTomlConf,
		dbProviders,
		controller.NewBlogController,
		controller.NewCateController,
		controller.NewHomeController)
	return nil, nil
}
