package Page

import (
	"gormuser/Common/dbase"
	"gormuser/Common/tool"

	"github.com/kataras/iris/v12"
)

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

//  127.0.0.1:84/Reply?token=xxx&rid=5&message=ccccccccccccc (√)
func Reply(ctx iris.Context) {
	sreid := ctx.URLParam("rid")
	msg := ctx.URLParam("message")
	uid := dbase.FromTokenGetID(ctx.URLParam("token"))

	if !tool.Isnumber(sreid) {
		ctx.Write([]byte("error"))
		return
	}
	reid := tool.String2Int(sreid)
	dbase.Reply(uid, reid, msg)
	return
}

//  127.0.0.1:84/PostMessage??token=xxx&title=KANG&message=ccccccccccccc (√)
func PostMessage(ctx iris.Context) {
	//Group_id
	//Message
	//title
	sid := ctx.URLParam("gid")
	uid := dbase.FromTokenGetID(ctx.URLParam("token"))
	if !tool.Isnumber(sid) {
		ctx.Write([]byte("error:0"))
		return
	}

	if dbase.PostMessage(uid, tool.String2Int(sid), ctx.URLParam("title"), ctx.URLParam("message")) {
		ctx.Write([]byte("ok"))
		return //Exit()
	}
	ctx.Write([]byte("error:1"))
	return
}

//  127.0.0.1:84/GetReply?reid=5&message=kkk&token= (×)
func GetReply(ctx iris.Context) {
	srid := ctx.URLParam("reid")
	uid := dbase.FromTokenGetID(ctx.URLParam("token"))
	if !tool.Isnumber(srid){
		ctx.Write([]byte("error:0"))
		return
	}
	if dbase.Reply(uid,tool.String2Int(srid),ctx.URLParam("message")){
		ctx.Write([]byte("ok"))
		return
	}
	ctx.Write([]byte("error:1"))

}
