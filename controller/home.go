package controller

import (
	"net/http"

	"github.com/Wuvist/decho/models"
	"github.com/Wuvist/decho/tpl"
	"github.com/Wuvist/decho/tpl/skins"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// HomeController handles homepage request
type HomeController struct {
	bloggers models.BloggerQuery
}

// NewHomeController return HomeController bind with echo engine
func NewHomeController(e *echo.Echo) (*HomeController, error) {
	ctrl := &HomeController{}
	e.GET("/", ctrl.show)

	return ctrl, nil
}

func (ctrl *HomeController) show(c echo.Context) error {
	objs, err := ctrl.bloggers(qm.Select("blogname", "id", "nick"),
		qm.Where("reveal = 1 and blogs > ?", 0),
		qm.OrderBy("last_post desc")).AllG()

	if err != nil {
		println(err.Error())
	}

	bloggers := make([]*tpl.Blogger, len(objs))
	for i := 0; i < len(objs); i++ {
		b := &tpl.Blogger{}
		obj := objs[i]
		b.Username = obj.ID
		b.BlogName = obj.Blogname
		b.Nick = obj.Nick.String
		bloggers[i] = b
	}

	page := skins.Home(bloggers)
	return c.HTML(http.StatusOK, page)
}
