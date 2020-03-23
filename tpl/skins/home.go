package skins

import (
	"bytes"
	"github.com/Wuvist/goblog/tpl"
	"github.com/sipin/gorazor/gorazor"
)

func Home(bloggers []*tpl.Blogger) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<html><head>\n<title>博客风</title>\n<link rel=\"stylesheet\" href=\"/Template/skin5.css\" type=\"text/css\"> \n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n</head>\n<body bgcolor=\"ffffff\" marginwidth=\"0\" marginheight=\"0\">\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tr bgcolor=\"#ffffff\">\n<td class=\"name\"  width=\"100%\" height=\"60\">博客风\n</td>\n</tr>\n</table>\n\n<table width=\"100%\" border=\"0\" cellspacing=\"0\" cellpadding=\"0\">\n<tr bgcolor=\"#990000\">\n<td class=\"description\" align=\"right\" width=\"100%\" height=\"30\">感谢您还记得博客风~</td>\n</tr>\n</table>\n\n<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\"><tr><td valign=\"top\">\n<div class=\"blog\">")
	for _, blogger := range bloggers {

		_buffer.WriteString("<div class=\"datetimefooter\"><a href=\"/")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Username))
		_buffer.WriteString("\" class=\"title\">")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.BlogName))
		_buffer.WriteString("</a>\n\t\t   ")
		_buffer.WriteString(gorazor.HTMLEscape(blogger.Nick))
		_buffer.WriteString("\n\t\t  </div><br>")
	}
	_buffer.WriteString("\n</div>\n</div>\n</TD>\n<TD bgColor=\"#e0e0b1\" height=\"500\" width=\"150\" valign=\"top\"><div class=\"side\">\n<div class=\"sidetop\">\n导航\n</div> \n  <a href=\"/\">博客风</a><BR>\n<br>\n<br>\n<br><br><br><div class=\"sidetop\">\n站点声明\n</div>\n<br /><br />\n自2003年上线至今（2018年），博客风已经运作了近15年。今年春节前服务器硬盘出现硬件故障，<a href=\"https://github.com/Wuvist/goblog\">经过努力</a>目前站点恢复了2017年年底的一个数据库备份；但仅提供静态访问；网志发表、评论功能关闭。\n<br /><br />\n您如果还想写博客，建议考虑<a href=\"http://www.jianshu.com\">简书</a>等网站。\n<br /><br />\n感谢您还记得博客风！\n<br /><br />\n此致，\n问天\n<br>\n<br>\n</div>\n</TD>\n</TR>\n</TABLE>\n</body>\n</html>")

	return _buffer.String()
}
