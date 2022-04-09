package Page

import (
	"gormuser/Common/User"
	"strings"

	"github.com/kataras/iris/v12"
)

func PageView(ctx iris.Context, name string) {
	if strings.Contains(name, ".") {
		if name[strings.LastIndex(name, "."):] != ".html" {
			ctx.View(name + ".html")
		} else {
			ctx.View(name)
		}
	} else {
		ctx.View(name + ".html")

	}

}
func PageError(ctx iris.Context) {
	ctx.Header("Authorization", "xxxxxx")
	ctx.View("error.html")
}

func Login(ctx iris.Context) {
	if User.IfLogin(ctx) {
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( '已登录!!'); window.location.href='../';</script>")
	} else {
		ctx.View("login.html")

	}
}
