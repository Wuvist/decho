//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"

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
	wire.Build(echo.New, wire.Struct(new(WebApp), "*"), loadTomlConf,
		dbProviders,
		controller.NewBlogController,
		controller.NewCateController,
		controller.NewHomeController,
		controller.NewStaticController)
	return nil, nil
}
