package main

import (
	"database/sql"

	"github.com/Wuvist/decho/conf"
	"github.com/Wuvist/decho/controller"
	"github.com/Wuvist/decho/models"
	"github.com/Wuvist/dechoconf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func main() {
	webApp, err := getWebApp()
	if err != nil {
		panic(err)
	}
	webApp.Run()
}

// WebApp represent a web application
type WebApp struct {
	*echo.Echo
	config *conf.App
	blog   *controller.BlogController
	cate   *controller.CateController
	home   *controller.HomeController
	static *controller.StaticController
}

// Run the web app
func (e *WebApp) Run() {
	db, err := sql.Open("mysql", e.config.Mysql+"?parseTime=true")
	if err != nil {
		panic(err)
	}
	boil.SetDB(db)
	e.Logger.Fatal(e.Start(e.config.Address))
}

func loadTomlConf() (*conf.App, error) {
	var conf conf.App
	if err := dechoconf.DecodeTomlFile("conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

// db providers
func getBloggerDB() models.BloggerQuery {
	return models.Bloggers
}

func getArticalDB() models.ArticlesQuery {
	return models.Articles
}

func getCategoryDB() models.UserdefinecategoryQuery {
	return models.Userdefinecategories
}

func getLinkDB() models.LinkQuery {
	return func(mods ...qm.QueryMod) models.LinkExecutor {
		return models.Links(mods...)
	}
}

func getCommentDB() models.CommentQuery {
	return models.Comments
}
