package main

import (
	"github.com/BurntSushi/toml"
	"github.com/Wuvist/decho/conf"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	InitEngine(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func newMsg() string {
	return "bingo"
}

func loadTomlConf() (*conf.Config, error) {
	var conf conf.Config
	if _, err := toml.DecodeFile("conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
