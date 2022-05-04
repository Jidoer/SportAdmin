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

func Message(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.ViewData("ListMessage", dbase.ListMessage(1)) //1：List ALL!
	ctx.View("message.html")
}
func HotMessage(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.ViewData("ListMessage", dbase.ListMessage(0)) //1：List HOT!
	ctx.View("message.html")
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
func DelMessage(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if tool.Isnumber(sid) {
		id := tool.String2Int(sid)
		if dbase.DelMessage(id) {
			ctx.HTML("<script language='javascript' type='text/javascript'> window.location.href='../message';</script>")
		} else {
			ctx.HTML("<script language='javascript' type='text/javascript'> window.location.href='../message';</script>")
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
	sex := tool.String2Int(ctx.URLParam("sex")) //ctx.URLParam("sex")
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
	if sex != 0 && sex != 1 {
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( '失败!!'); window.location.href='../admin';</script>")
		return
	}

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
	if dbase.EditUser(id, newuser) {
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../admin';</script>")

		return
	}
	ctx.HTML("<script language='javascript' type='text/javascript'> alert( '失败!!'); window.location.href='../admin';</script>")
	return
}

func Hot(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.View("hot.html")
}

func EditMsg(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error: ID Is Not Num!"))
		return
	}
	title := ctx.URLParam("title")
	msg := ctx.URLParam("message")
	if title == "" || msg == "" {
		oldMsg := dbase.GetMessageInfo(tool.String2Int(sid))
		if title == "" {
			title = oldMsg.Title
		}
		if msg == "" {
			msg = oldMsg.Message
		}
	}
	dbase.EditMsg(tool.String2Int(sid), dbase.Discuss{
		Title:   title,
		Message: msg,
	})
}

func DelReply(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error"))
		return
	}
	if dbase.DelReply(tool.String2Int(sid)) {
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../message';</script>")
		return
	}
	ctx.Write([]byte("[error]"))
}

func ToHot(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error: ID Is Not Num!"))
		return
	}
	if dbase.ToHot(tool.String2Int(sid)) {
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../message';</script>")

	} else {
		ctx.Write([]byte("error!"))
	}

}

func NoToHot(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error: ID Is Not Num!"))
		return
	}
	if dbase.NoToHot(tool.String2Int(sid)) {
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../message';</script>")

	} else {
		ctx.Write([]byte("error!"))
	}
}

func FeedBack(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.ViewData("ListFeedBacks", dbase.ListFeedBacks()) //ALL!
	ctx.View("feedback.html")

	//ListFeedBacks
}

func ReplyView(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	ctx.ViewData("ReplyView", dbase.ReplyView(tool.String2Int(ctx.URLParam("id")))) //ALL!
	ctx.View("reply.html")

	//ListFeedBacks
}


func EditReply(ctx iris.Context){
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	id := ctx.URLParam("id")
	msg := ctx.URLParam("message")
	if dbase.EditReply(tool.String2Int(id),msg){
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../feedback';</script>")
		return
	}
	ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../feedback';</script>")

}

func DelFeedBack(ctx iris.Context) {
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error"))
		return
	}
	if dbase.DelFeedBack(tool.String2Int(sid)) {
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../feedback';</script>")
		return
	}
	ctx.Write([]byte("[error]"))
}


func EditFeedback(ctx iris.Context){
	if !User.IsAdmin(ctx) {
		ctx.View("error.html")
		return
	}
	id := ctx.URLParam("id")
	msg := ctx.URLParam("message")
	if dbase.EditFeedback(tool.String2Int(id),msg){
		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../feedback';</script>")
		return
	}
	ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../feedback';</script>")

}


func AddMessage(ctx iris.Context){
	if !User.IsAdmin(ctx){
		ctx.View("error.html")
		return
	}
	if dbase.AddMeassage(ctx.URLParam("title"),ctx.URLParam("message")){

		ctx.HTML("<script language='javascript' type='text/javascript'>window.location.href='../message';</script>")

		return
	}
}
