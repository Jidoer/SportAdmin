package Page

import (
	"gormuser/Common/User"
	"gormuser/Common/dbase"
	"gormuser/Common/tool"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

/*
type ReplayMsg struct {
	ID         string `json:"id"`
	HeadImg    string `json:"headImg"`
	UserName   string `json:"userName"`
	UserID     string `json:"userId"`
	Content    string `json:"content"`
	CreatTime  string `json:"creatTime"`
	LikeCount  string `json:"likeCount"`
	IsLike     string `json:"isLike"`
	TotalCount string `json:"totalCount"`
}
*/
// 127.0.0.1:84/GetList?group=0 1 2 (√)
func GetList(ctx iris.Context) {
	c := dbase.ListDiscuss(tool.String2Int(ctx.URLParam("group")))
	ctx.Write([]byte(tool.InterfaceToJson(c)))
}

// 127.0.0.1:84/GetMessageInfo?id=5 (√)
func GetMessageInfo(ctx iris.Context) {
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error"))
		return
	}
	Info := dbase.GetMessageInfo(tool.String2Int(sid))
	ctx.Write([]byte(tool.InterfaceToJson(Info)))
	return
}

//  127.0.0.1:84/PostMessage??token=xxx&title=KANG&message=ccccccccccccc &gid=0 1 2  (√)
func PostMessage(ctx iris.Context) {

	if !dbase.CKToken(ctx.URLParam("token")){
		log.Println("Token Error!")
		ctx.Write([]byte("error:0"))
		return
	}

	uid := dbase.FromTokenGetID(ctx.URLParam("token"))

	if dbase.PostMessage(uid, 1, ctx.URLParam("title"), ctx.URLParam("message")) {
		log.Printf("提交->标题:%s内容%s\n", ctx.URLParam("title"), ctx.URLParam("message"))
		ctx.Write([]byte("ok"))
		return //Exit()
	}
	ctx.Write([]byte("error:1"))
	return
}

//  127.0.0.1:84/Reply?token=xxx&rid=5&message=ccccccccccccc (√)
func Reply(ctx iris.Context) {
	sreid := ctx.URLParam("rid")
	msg := ctx.URLParam("message")

	if !dbase.CKToken(ctx.URLParam("token")){
		log.Println("Token Error!")
		ctx.Write([]byte("error:0"))
		return
	}

	uid := dbase.FromTokenGetID(ctx.URLParam("token"))
	if uid == -1 {
		ctx.Write([]byte("error"))
		return
	}
	if !tool.Isnumber(sreid) {
		ctx.Write([]byte("error"))
		return
	}
	reid := tool.String2Int(sreid)
	dbase.Reply(uid, reid, msg)
	log.Printf("回复->：%s内容：%s\n", ctx.URLParam("rid"), ctx.URLParam("message"))
	return
}

/*
//  127.0.0.1:84/GetReply?reid=5&message=kkk&token= (×)
func GetReply(ctx iris.Context) {
	srid := ctx.URLParam("reid")
	uid := dbase.FromTokenGetID(ctx.URLParam("token"))
	if !tool.Isnumber(srid) {
		ctx.Write([]byte("error:0"))
		return
	}
	if dbase.Reply(uid, tool.String2Int(srid), ctx.URLParam("message")) {
		ctx.Write([]byte("ok"))
		return
	}
	ctx.Write([]byte("error:1"))
}
*/

func AppLogin(ctx iris.Context) {
	username := ctx.URLParam("user")
	pw := ctx.URLParam("pw")
	log.Print("/AppLogin")
	result, info := User.Login(username, pw)
	if result == "yes" {
		//生成Token
		token := tool.GetStringMd5(time.Now().String() + strconv.Itoa(rand.Intn(10000)+1000))
		if dbase.SetToken(info, token) {
			ctx.SetCookieKV("token", token)
			log.Println("APP登陆成功")
			ctx.Write([]byte("<ok>[" + token + "]"))
			return
		} else {
			log.Println("error 1")
			ctx.Write([]byte("[error]"))
			return
		}

	} else {
		log.Println("error 2")
		ctx.Write([]byte("[error]"))
		return
	}
}

func APPCK(ctx iris.Context) {
	token := ctx.URLParam("token")
	if dbase.CKToken(token) {
		ctx.Write([]byte("[ok]"))
		return
	}
	ctx.Write([]byte("error"))
	return
}

func AppReg(ctx iris.Context) {
	//username := ctx.PostValue("user")
	username := ctx.URLParam("user")
	pw := ctx.URLParam("pw")
	email := ctx.URLParam("email")
	phone := ctx.URLParam("phone")
	var user dbase.UserInfo
	user = dbase.UserInfo{
		Username: username,
		Password: pw,
		Money:    0,
		Vip:      0,
		Sex:      0,
		Phone:    phone,
		Email:    email,
		Loginip:  ctx.RemoteAddr(),
	}
	result := User.Reg(user)
	if result == "yes" {
		log.Println("APP注册成功：" + username)
		ctx.Write([]byte("[ok]"))
	} else if result == "re" {
		ctx.Write([]byte("[re]"))
		log.Println("用户名重复：" + username)
	} else {
		log.Println("APP注册失败：" + username)
		//print(time.April.String() + "Loginerror: [TIME]" +username + "password:" + pw+"code:"+ result +"\n")
		ctx.Write([]byte("[no]"))
	}

}

/////////////////////////////////////////////////////////////////////

func PostFeedback(ctx iris.Context) {
	if dbase.PostFeedback(ctx.URLParam("email"), ctx.URLParam("msg"), tool.String2Int(ctx.URLParam("type"))) {
		ctx.Write([]byte("[ok]"))
		log.Println("提交反馈: " + ctx.URLParam("msg"))
		return
	}
	ctx.Write([]byte("[no]"))
	return
}

func GetReplyList(ctx iris.Context) {
	sid := ctx.URLParam("id")
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error"))
		return
	}
	ctx.Write([]byte(tool.InterfaceToJson(dbase.ListReply(tool.String2Int(sid)))))
	//ctx.ViewData("ListReply",tool.InterfaceToJson(dbase.ListReply(tool.String2Int(sid))))
	//ctx.View("listreply.html")
	return
}

//(x)未添加 SendEmail
func SendEmail(ctx iris.Context) {
	email := ctx.URLParam("email")
	message := ctx.URLParam("message")
	log.Println("发送邮件:" + email + "内容：" + message)
}

func AddMsg(ctx iris.Context) {
	if dbase.PostFeedback(ctx.URLParam("email"), ctx.URLParam("msg"), tool.String2Int(ctx.URLParam("type"))) {
		ctx.Write([]byte("[ok]"))
		log.Println("提交反馈: " + ctx.URLParam("msg"))
		return
	}
	ctx.Write([]byte("[no]"))
	return
}

