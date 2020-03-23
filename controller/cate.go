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

// CateController handles blog category requests
type CateController struct {
	bloggers models.BloggerQuery
	cates    models.UserdefinecategoryQuery
	vm       *tpl.ViewModel
}

// NewCateController return CateController bind with echo engine
func NewCateController(e *echo.Echo, bloggers models.BloggerQuery,
	cates models.UserdefinecategoryQuery, vm *tpl.ViewModel) (*CateController, error) {
	cate := &CateController{
		bloggers,
		cates,
		vm,
	}
	e.GET("/cate.go", cate.show)

	return cate, nil
}

func (ctrl *CateController) show(c echo.Context) error {
	blogerUsername := strings.ToLower(c.QueryParam("blogger"))
	cateID := c.QueryParam("cate_id")
	cateData, _ := ctrl.cates(qm.Where("`index` = ?", cateID)).OneG()
	if cateData == nil {
		return c.String(http.StatusNotFound, "找不到网志分类")
	}

	bloggerData, _ := ctrl.bloggers(qm.Where("`index` = ?", cateData.Blogger)).OneG()
	if bloggerData == nil {
		return c.String(http.StatusNotFound, "找不到博客")
	}

	blogger := ctrl.vm.NewBloggerFromDb(bloggerData)
	if strings.ToLower(blogger.Username) != blogerUsername {
		return c.String(http.StatusNotFound, "找不到分类")
	}

	cate := &tpl.Cate{}
	cate.CateID = cateData.Index
	cate.CateName = cateData.Cate

	blogs := ctrl.vm.GetBlogSummariesFromCate(cate.CateID)

	page := skins.Skin5_UserCate(blogger, cate, blogs)

	return c.HTML(http.StatusOK, page)
}
