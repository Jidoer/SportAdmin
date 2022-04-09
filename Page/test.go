package Page

import (
	"gormuser/Common/dbase"
	"gormuser/Common/tool"
	"log"

	"github.com/kataras/iris/v12"
)

//Jast Demo Test API !!!!
func AddUser2(ctx iris.Context) {
	/*
		if !User.IsAdmin(ctx) {
			ctx.View("error.html")
			return
		}
	*/
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	sex := 0 //ctx.URLParam("sex")
	money := ctx.URLParam("money")
	vip := ctx.URLParam("vip")
	phone := ctx.URLParam("phone")
	email := ctx.URLParam("email")
	var newuser dbase.UserInfo
	newuser = dbase.UserInfo{
		Username: username,
		Password: password,
		Sex:      sex,
		Money:    tool.String2Int(money),
		Vip:      tool.String2Int(vip),
		Phone:    phone,
		Email:    email,
		Loginip:  "null", /*ctx.RemoteAddr()*/
	}
	log.Println("Add User 👇")
	log.Println(newuser)
	if dbase.AddUser(newuser) {
		ctx.HTML("<script language='javascript' type='text/javascript'> window.location.href='../admin';</script>")
	}
	ctx.HTML("<script language='javascript' type='text/javascript'> alert( '失败!!'); window.location.href='../admin';</script>")

	return
}
