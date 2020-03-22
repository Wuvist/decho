package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Wuvist/decho/conf"
	"github.com/Wuvist/decho/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	app, err := getApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}

func newMsg() string {
	return "bingo"
}

type webApp struct {
	*echo.Echo
	config *conf.Config
	blog   *controller.BlogController
}

func (e *webApp) Run() {
	e.Logger.Fatal(e.Start(":1323"))
}

func newEcho() *echo.Echo {
	return echo.New()
}

func newWebApp(config *conf.Config, e *echo.Echo, blog *controller.BlogController) *webApp {
	return &webApp{
		e,
		config,
		blog,
	}
}

func loadTomlConf() (*conf.Config, error) {
	var conf conf.Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
