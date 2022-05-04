package main

import (
	"gormuser/Page"
	pageerror "gormuser/PageError"

	"github.com/kataras/iris/v12"
)

func main() {
	admindir := "/admin"
	app := iris.New()
	app.Logger().SetLevel("debug")

	//tmpl := iris.HTML("./Template", ".html")
	tmpl := iris.Django("./Template", ".html")

	tmpl.Reload(true)
	// 设置页面的函数
	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings, " + s + "!"
	})
	app.RegisterView(tmpl)

	app.OnErrorCode(iris.StatusNotFound, pageerror.NotFound)

	//INDEX
	app.Get("/", func(ctx iris.Context) { Page.Admin(ctx) })
	app.Get("/account", func(ctx iris.Context) { Page.Admin(ctx) })
	app.Get("/account.html", func(ctx iris.Context) { Page.Admin(ctx) })
	app.Get(admindir+"/", func(ctx iris.Context) { Page.Admin(ctx) })
	app.Get("/login", func(ctx iris.Context) { Page.Login(ctx) })
	app.Get("/login", func(ctx iris.Context) { Page.Login(ctx) })
	app.Get("/message", func(ctx iris.Context) { Page.Message(ctx) })
	app.Get("/hot", func(ctx iris.Context) { Page.HotMessage(ctx) })
	app.Get("/feedback", func(ctx iris.Context) { Page.FeedBack(ctx) })
	app.Get("/sendemail",func(ctx iris.Context) {Page.SendEmail(ctx)})
	app.Get("/reply",func(ctx iris.Context) {Page.ReplyView(ctx)})


	//API
	//ListDiscuss
	app.Get("/GetList", func(ctx iris.Context) { Page.GetList(ctx) })
	app.Get("/GetMessageInfo", func(ctx iris.Context) { Page.GetMessageInfo(ctx) })
	app.Get("/Reply", func(ctx iris.Context) { Page.Reply(ctx) })
	app.Get("/PostMessage", func(ctx iris.Context) { Page.PostMessage(ctx) })
	app.Get("/AppLogin", func(ctx iris.Context) { Page.AppLogin(ctx) })
	app.Get("/GetReplyList", func(ctx iris.Context) { Page.GetReplyList(ctx) })
	app.Get("/PostFeedback", func(ctx iris.Context) { Page.PostFeedback(ctx) })
	app.Get("/editmsg", func(ctx iris.Context) { Page.EditMsg(ctx) })
	app.Get("/editreply", func(ctx iris.Context) { Page.EditReply(ctx) })
	app.Get("/tohot", func(ctx iris.Context) { Page.ToHot(ctx) })
	app.Get("/notohot", func(ctx iris.Context) { Page.NoToHot(ctx) })
	app.Get("/delfeedback", func(ctx iris.Context) { Page.DelFeedBack(ctx) })
	app.Get("/editfeedback", func(ctx iris.Context) { Page.EditFeedback(ctx) })
	app.Get("/addmsg", func(ctx iris.Context) { Page.AddMessage(ctx) })



	//PostApi
	app.Get("/appreg", func(ctx iris.Context) { Page.AppReg(ctx) })

	//Admin
	app.Get("/adduser", func(ctx iris.Context) { Page.AddUser(ctx) })
	app.Get("/deluser", func(ctx iris.Context) { Page.DellUser(ctx) })
	app.Get("/edituser", func(ctx iris.Context) { Page.EditUser(ctx) })
	app.Get("/delmessage", func(ctx iris.Context) { Page.DelMessage(ctx) })
	app.Get("/delreply", func(ctx iris.Context) { Page.DelReply(ctx) })
	

	app.Post("/regapi", func(ctx iris.Context) { Page.RegApi(ctx) })
	app.Post("/loginapi", func(ctx iris.Context) { Page.LoginApi(ctx) })
	app.Get("/run", func(ctx iris.Context) { Page.Run(ctx) })

	//testadduser Just Test API !!!!
	//app.Get("/testadduser", func(ctx iris.Context) { Page.AddUser2(ctx) })

	//-----new------
	app.Post("/edit", func(ctx iris.Context) { Page.RegApi(ctx) })
	app.Get("/del", func(ctx iris.Context) { Page.Del(ctx) })
	app.Get("api/{path:path}", func(ctx iris.Context) {
		apiCall := ctx.Params().Get("path")
		app.Logger().Info(apiCall)
	})

	app.HandleDir("/js", "./Template/js")
	app.HandleDir("/css", "./Template/css")
	app.HandleDir("/fonts", "./Template/fonts")
	app.HandleDir("/images", "./Template/images")
	app.HandleDir("/vendors", "./Template/vendors")

	app.Run(iris.Addr("0.0.0.0:84"))
}
