package tpl

import (
	"github.com/Wuvist/decho/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

var dateFormat = "2006-01-02 15:04:05"

// BlogSummary is the struture for displaying BlogSummary
type BlogSummary struct {
	ID           int
	Title        string
	Adddate      string
	CommentCount int
}

// Blog is the struture for displaying a blog
type Blog struct {
	Title    string
	CateName string
	CateID   int
	Content  string
	Tags     string
	Adddate  string
}

// ViewModel is helper struct for rendering views
type ViewModel struct {
	Articles models.ArticlesQuery
	Comments models.CommentQuery
	Cates    models.UserdefinecategoryQuery
	Links    models.LinkQuery
}

// GetBlogSummariesFromCate get a blogs of a blog category
func (vm *ViewModel) GetBlogSummariesFromCate(cateID int) []*BlogSummary {
	articleData, _ := vm.Articles(
		qm.Select("index", "title", "add_date", "Comment"),
		qm.Where("`cate` = ?", cateID),
		qm.OrderBy("`index` desc")).AllG()

	blogs := make([]*BlogSummary, len(articleData))
	for i := 0; i < len(articleData); i++ {
		b := &BlogSummary{}
		obj := articleData[i]

		b.ID = obj.Index
		b.CommentCount = obj.Comment
		b.Title = obj.Title.String
		b.Adddate = obj.AddDate.Time.Format(dateFormat)
		blogs[i] = b
	}

	return blogs
}

// GetBlogSummariesFromBlogger get top 20 blogs of a blogger
func (vm *ViewModel) GetBlogSummariesFromBlogger(bloggerID int) []*BlogSummary {
	articleData, _ := vm.Articles(
		qm.Select("index", "title", "add_date", "Comment"),
		qm.Where("`blogger` = ?", bloggerID),
		qm.OrderBy("`index` desc"),
		qm.Limit(20)).AllG()

	blogs := make([]*BlogSummary, len(articleData))
	for i := 0; i < len(articleData); i++ {
		b := &BlogSummary{}
		obj := articleData[i]

		b.ID = obj.Index
		b.CommentCount = obj.Comment
		b.Title = obj.Title.String
		b.Adddate = obj.AddDate.Time.Format(dateFormat)
		blogs[i] = b
	}

	return blogs
}

// GetBlogComments get all comments of a blog
func (vm *ViewModel) GetBlogComments(blogID int) []*Comment {
	objs, _ := vm.Comments(
		qm.Where("`article` = ?", blogID),
		qm.OrderBy("`index` desc")).AllG()

	comments := make([]*Comment, len(objs))
	for i := 0; i < len(objs); i++ {
		c := &Comment{}
		obj := objs[i]
		c.Author = obj.Author.String
		c.Content = obj.Content
		c.Adddate = obj.AddDate.Format(dateFormat)
		comments[i] = c
	}
	return comments
}

// NewBlogFromDb create blog struct from db model
func (vm *ViewModel) NewBlogFromDb(data *models.Article) *Blog {
	b := &Blog{}
	b.Title = data.Title.String
	if b.Title == "" {
		b.Title = "无题"
	}
	b.Content = data.Content.String
	b.Adddate = data.AddDate.Time.Format(dateFormat)

	cateData, _ := vm.Cates(qm.Where("`index` = ?", data.Cate.Int)).OneG()
	if cateData != nil {
		b.CateID = cateData.Index
		b.CateName = cateData.Cate
	}

	return b
}

// Cate is blogger's own blog categories
type Cate struct {
	CateName  string
	CateID    int
	BlogCount int
}

// Link is blogger's own links
type Link struct {
	URL  string
	Link string
}

// Blogger is the struture for blogger info
type Blogger struct {
	Username     string
	BlogName     string
	Info         string
	Nick         string
	VisitorCount int
	Cates        []*Cate
	Links        []*Link
}

func (vm *ViewModel) getBloggerCates(bloggerID int) []*Cate {
	objs, _ := vm.Cates(qm.Where("`blogger` = ?", bloggerID)).AllG()
	cates := make([]*Cate, len(objs))
	for i := 0; i < len(objs); i++ {
		cate := &Cate{}
		obj := objs[i]
		cate.CateName = obj.Cate
		cate.CateID = obj.Index
		cate.BlogCount = obj.Files
		cates[i] = cate

	}
	return cates
}

func (vm *ViewModel) getBloggerLinks(bloggerID int) []*Link {
	objs, _ := vm.Links(qm.Where("`blogger` = ? and reveal = 1",
		bloggerID)).AllG()
	links := make([]*Link, len(objs))
	for i := 0; i < len(objs); i++ {
		link := &Link{}
		obj := objs[i]
		link.Link = obj.Link
		link.URL = obj.URL
		links[i] = link

	}
	return links
}

// NewBloggerFromDb create blogger struct from db model
func (vm *ViewModel) NewBloggerFromDb(data *models.Blogger) *Blogger {
	b := &Blogger{}
	b.Username = data.ID
	b.Nick = data.Nick.String
	b.Info = data.Intro.String
	b.BlogName = data.Blogname
	b.VisitorCount = data.Visitor

	b.Cates = vm.getBloggerCates(data.Index)
	b.Links = vm.getBloggerLinks(data.Index)

	return b
}

// Comment is the struture for displaying a comment
type Comment struct {
	Author  string
	Adddate string
	Content string
}
