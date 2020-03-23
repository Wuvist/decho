//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Wuvist/decho/controller"
	"github.com/Wuvist/decho/tpl"
)

var dbProviders = wire.NewSet(
	wire.Struct(new(tpl.ViewModel), "*"),
	getBloggerDB,
	getArticalDB,
	getCategoryDB,
	getCommentDB,
	getLinkDB,
)

func getWebApp() (*WebApp, error) {
	wire.Build(newEcho, wire.Struct(new(WebApp), "*"), loadTomlConf,
		dbProviders,
		controller.NewBlogController,
		controller.NewCateController,
		controller.NewHomeController)
	return nil, nil
}
