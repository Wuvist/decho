//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Wuvist/decho/controller"
)

func getApp() (*webApp, error) {
	wire.Build(newEcho, newWebApp, loadTomlConf, controller.NewBlogController, newMsg)
	return nil, nil
}
