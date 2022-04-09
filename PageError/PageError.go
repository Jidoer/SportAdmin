package pageerror

import "github.com/kataras/iris/v12"

func NotFound(ctx iris.Context) {
	ctx.View("404.html")
}
