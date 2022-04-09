package Page

import (
	"gormuser/Common/User"
	"gormuser/Common/dbase"
	"gormuser/Common/tool"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
)

func Admin(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.ViewData("ListUser", dbase.ListUser())
	ctx.View("account.html")
}

func AddUser(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
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

func DellUser(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if tool.Isnumber(sid) {
		id := tool.String2Int(sid)
		if dbase.DelUser(id) {
			ctx.HTML("<script language='javascript' type='text/javascript'> window.location.href='../admin';</script>")
		} else {
			ctx.HTML("<script language='javascript' type='text/javascript'> window.location.href='../admin';</script>")
		}
	}
	return
}

func EditUser(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		log.Println("NO ADMIN")
		ctx.View("error.html")
		return
	}
	username := ctx.URLParam("username")
	password := ctx.URLParam("password")
	sex := 0 //ctx.URLParam("sex")
	money := ctx.URLParam("money")
	vip := ctx.URLParam("vip")
	phone := ctx.URLParam("phone")
	email := ctx.URLParam("email")
	sid := ctx.URLParam("id")
	/*
		if(tool.Isnumber(sid)){
			ctx.View("error.html")
			return
		}
	*/
	log.Println(sid)
	id := tool.String2Int(sid)

	//GetOldUser
	olduser, er := dbase.GetInfoFromID(id)
	if er != "yes" {
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( '失败!!'); window.location.href='../admin';</script>")
		return
	}

	if password == "" {
		password = olduser.Password
	}
	if username == "" {
		username = olduser.Username
	}
	if phone == "" {
		phone = olduser.Phone
	}
	if money == "" {
		money = strconv.Itoa(olduser.Money)
	}
	if email == "" {
		email = olduser.Email
	}

	var newuser dbase.UserInfo
	newuser = dbase.UserInfo{
		Username: username,
		Password: password,
		Sex:      sex,
		Money:    tool.String2Int(money),
		Vip:      tool.String2Int(vip),
		Phone:    phone,
		Email:    email,
	}
	if dbase.EditUser(id, newuser){
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../admin';</script>")

		return
	}
	ctx.HTML("<script language='javascript' type='text/javascript'> alert( '失败!!'); window.location.href='../admin';</script>")
	return
}
