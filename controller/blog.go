package controller

import (
	"net/http"
	"strings"

	"github.com/Wuvist/decho/models"
	"github.com/Wuvist/decho/tpl"
	"github.com/Wuvist/decho/tpl/skins"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// BlogController defines controller for blog module
type BlogController struct {
	bloggers models.BloggerQuery
	articles models.ArticlesQuery
	vm       *tpl.ViewModel
}

func (ctrl *BlogController) show(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
	blogID := c.QueryParam("article_id")
	blogData, err := ctrl.articles(qm.Where("`index` = ?", blogID)).OneG()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	bloggerData, _ := ctrl.bloggers(qm.Where("`index` = ?", blogData.Blogger)).OneG()
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := ctrl.vm.NewBloggerFromDb(bloggerData)
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到网志")
	}

	blog := ctrl.vm.NewBlogFromDb(blogData)
	comments := ctrl.vm.GetBlogComments(blogData.Index)

	page := skins.Skin5_comment(blogger, blog, comments)

	return c.HTML(http.StatusOK, page)
}

func (ctrl *BlogController) showBlogger(c echo.Context) error {
	blogerUsername := c.QueryParam("blogger")

	bloggerData, err := ctrl.bloggers(qm.Where("id = ?", blogerUsername)).OneG()
	if err != nil {
		println(err.Error())
	}
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := ctrl.vm.NewBloggerFromDb(bloggerData)

	blogs := ctrl.vm.GetBlogSummariesFromBlogger(bloggerData.Index)

	page := skins.Skin5_default(blogger, blogs)

	return c.HTML(http.StatusOK, page)
}

// NewBlogController return BlogController bind with echo engine
func NewBlogController(e *echo.Echo, bloggers models.BloggerQuery,
	articles models.ArticlesQuery, vm *tpl.ViewModel) (*BlogController, error) {
	blog := &BlogController{
		bloggers,
		articles,
		vm,
	}
	e.GET("/blog.go", blog.show)
	e.GET("/blogger.go", blog.showBlogger)

	return blog, nil
}
