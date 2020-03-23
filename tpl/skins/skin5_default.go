package skins

import (
	"bytes"
	"github.com/Wuvist/goblog/tpl"
	"github.com/sipin/gorazor/gorazor"
)

func Skin5_default(blogger *tpl.Blogger, blogs []*tpl.BlogSummary) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<html><head>\n<title>")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.BlogName))
	_buffer.WriteString("</title>\n<link rel=\"stylesheet\" href=\"/Template/skin5.css\" type=\"text/css\"> \n<script type=\"text/javascript\" src=\"/icon.js\"></script>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n</head>\n<body bgcolor=\"ffffff\" marginwidth=\"0\" marginheight=\"0\">\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tr bgcolor=\"#ffffff\">\n<td class=\"name\"  width=\"100%\" height=\"60\">")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.BlogName))
	_buffer.WriteString("\n</td>\n</tr>\n</table>\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tr bgcolor=\"#990000\">\n<td class=\"description\" align=\"right\" width=\"100%\" height=\"30\">")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.Info))
	_buffer.WriteString("</td>\n</tr>\n</table>\n\n<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\"><tr><td valign=\"top\">\n<div class=\"blog\">")
	for _, blog := range blogs {

		_buffer.WriteString("<div class=\"datetimefooter\"><a href=\"/")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Username))
		_buffer.WriteString("/")
		_buffer.WriteString(gorazor.HTMLEscape((blog.ID)))
		_buffer.WriteString(".shtml\" class=\"title\">")
		_buffer.WriteString(gorazor.HTMLEscape(blog.Title))
		_buffer.WriteString("</a>\n\t\t   ")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Nick))
		_buffer.WriteString(" @ ")
		_buffer.WriteString(gorazor.HTMLEscape(blog.Adddate))
		_buffer.WriteString("\n\t\t <a href=\"/")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Username))
		_buffer.WriteString("/")
		_buffer.WriteString(gorazor.HTMLEscape((blog.ID)))
		_buffer.WriteString(".shtml\">评论:")
		_buffer.WriteString(gorazor.HTMLEscape(blog.CommentCount))
		_buffer.WriteString("</a></div><br>")
	}
	_buffer.WriteString("\n</div>\n</div>\n</TD>\n<TD bgColor=\"#e0e0b1\" height=\"500\" width=\"150\" valign=\"top\"><div class=\"side\">\n<div class=\"sidetop\">\n 导航\n </div> \n  <a href=\"/\">博客风</a><BR>\n  <a href=\"/")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.Username))
	_buffer.WriteString("/\">")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.BlogName))
	_buffer.WriteString("首页</a><BR>\n  <a>联系</a><BR>\n<br>\n\n<div class=\"sidetop\"><br>\n个人档案\n</div>\n<div align='center'><br><img src='//storage.googleapis.com/blogwind/images/userpic/")
	_buffer.WriteString(gorazor.HTMLEscape((blogger.Username)))
	_buffer.WriteString(".jpg'></div>")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.Info))
	_buffer.WriteString("\n<br>\n<br><br><br><div class=\"sidetop\">\n网志分类\n</div>")
	for _, cate := range blogger.Cates {

		_buffer.WriteString("<a href = \"/")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Username))
		_buffer.WriteString("/cate")
		_buffer.WriteString(gorazor.HTMLEscape((cate.CateID)))
		_buffer.WriteString(".shtml\" >")
		_buffer.WriteString(gorazor.HTMLEscape(cate.CateName))
		_buffer.WriteString("</a><span>(</span>")
		_buffer.WriteString(gorazor.HTMLEscape(cate.BlogCount))
		_buffer.WriteString("<span>)</span><br>")
	}
	_buffer.WriteString("\n<br><br><br><div class=\"sidetop\">\n网志存档\n</div>\n<!-- <asp:Repeater id=\"archive\" runat=\"server\">\n\t<ItemTemplate>\n\t\t<a href = \"/<%# uid & \"/\" & Container.DataItem(\"year\") %>/<%# Container.DataItem(\"month\") %>/\" ><%# Container.DataItem(\"year\") %>年<%# Container.DataItem(\"month\") %>月</a><br>\n\t</ItemTemplate>\n</asp:Repeater> -->\n\n<br><br><br><div class=\"sidetop\">\n个人链接\n</div>")
	for _, link := range blogger.Links {

		_buffer.WriteString("<a href=\"")
		_buffer.WriteString(gorazor.HTMLEscape(link.URL))
		_buffer.WriteString("\" target=\"_blank\">")
		_buffer.WriteString(gorazor.HTMLEscape(link.Link))
		_buffer.WriteString("</a><br>")
	}
	_buffer.WriteString("\n<p class=\"nav\">累计浏览: ")
	_buffer.WriteString(gorazor.HTMLEscape(blogger.VisitorCount))
	_buffer.WriteString("</p>\n<br>\n<br>\n</div>\n</TD>\n</TR>\n</TABLE>\n</body>\n</html>")

	return _buffer.String()
}
