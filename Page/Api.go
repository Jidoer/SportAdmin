package Page

import (
	//"crypto/md5"
	//"encoding/base64"
	"gormuser/Common/dbase"
	"gormuser/Common/tool"
	//"gormuser/Common/tool/file"
	//"gormuser/Common/tool/mydes"

	"github.com/kataras/iris/v12"
	"gormuser/Common/User"
	pageerror "gormuser/PageError"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func LoginApi(ctx iris.Context) {
	/*
	*
	 */
	username := ctx.PostValue("user")
	pw := ctx.PostValue("pw")
	log.Print("/Login")
	result, info := User.Login(username, pw)
	if result == "yes" {

		//生成Token
		token := tool.GetStringMd5(time.Now().String() + strconv.Itoa(rand.Intn(10000)+1000))
		if dbase.SetToken(info, token) {
			ctx.SetCookieKV("token", token)
			ctx.HTML("<script language='javascript' type='text/javascript'> alert( 'Successful!'); window.location.href='admin';</script>")

		} else {
			ctx.HTML("<script language='javascript' type='text/javascript'> alert( 'Set Token Error!'); window.location.href='/login';</script>")

		}
		/* 更新Token方式
		randnumber := rand.Intn(100)
		//
		key := base64.URLEncoding.EncodeToString([]byte(string(rune(randnumber)))) //zan shi
		kusername := base64.URLEncoding.EncodeToString([]byte(username))
		cookieskey, err := mydes.Encrypt(key, []byte("2fa6c1e9"))
		//
		if err != nil {
			log.Println(err)
		}
		fileuser, err := mydes.Encrypt(kusername, []byte("2fa6c1e9"))
		if err != nil {
			log.Println(err)
		}
		file.Create("Login/cookies/"+fileuser, cookieskey) //Base64-URL
		//ctx.SetCookieKV("user",kusername) //Base64-URL
		ctx.SetCookieKV("username", username) //前端用户名不加密
		ctx.SetCookieKV("key", cookieskey)
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( 'Successful!'); window.location.href='admin';</script>")
		//systeminfo.MakeSuccessInfoAlert("登录成功","登录成功!","200",ctx)
		//Index(ctx)
		ctx.ViewData("page", "index")
		ctx.ViewData("iflogin", true)
		ctx.View("account.html")
		*/

	} else {
		print(time.April.String() + "Loginerror: [TIME]" + username + "password:" + pw + "code:" + result + "\n")
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( '用户名或密码错误!'); window.location.href='/login';</script>")
	}
	//ctx.HTML("404")
}

func RegApi(ctx iris.Context) {
	/*
	*
	 */
	username := ctx.PostValue("user")
	pw := ctx.PostValue("pw")
	email := ctx.PostValue("email")
	phone := ctx.PostValue("phone")
	var user dbase.UserInfo
	user = dbase.UserInfo{
		Username: username,
		Password: pw,
		Money: 0,
		Vip: 0,
		Sex: 0,
		Phone: phone,
		Email: email,
		Loginip: ctx.RemoteAddr(),
	}
	result := User.Reg(user)
	if result == "yes" {
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( 'Successful!'); window.location.href='admin';</script>")

	} else if result == "re" {
		ctx.HTML("<script language='javascript' type='text/javascript'>alert( '用户名重复!' ); window.location.href='reg/';</script>")

	} else {
		//print(time.April.String() + "Loginerror: [TIME]" +username + "password:" + pw+"code:"+ result +"\n")

		ctx.HTML("<script language='javascript' type='text/javascript'>alert( 'Error Code : 201' ); window.location.href='reg/';</script>")
	}

	//ctx.HTML("404")
}

func Run(ctx iris.Context) {
	/*****
	 *Api For some activities
	 */
	exec := string(ctx.URLParam("exec"))
	if exec == "loginout" {
		//Login Out o_0
		ctx.SetCookieKV("token", "")
		ctx.HTML("<script language='javascript' type='text/javascript'> alert( '您已成功退出登录!'); window.location.href='/';</script>")
	} else if exec == "userinfo" {
		//ctx.HTML("Error Code: 103!")
		result := User.GetInfoFromCTX(ctx)
		if result == nil {
			ctx.HTML("Error Code 100")
		} else {
			ctx.HTML(tool.InterfaceToJson(result)) //输出User Info Json
		}
	} else if exec == "getserver" {

	} else {
		//---->
	}
}

func Del(ctx iris.Context) {
	//先鉴权
	if User.IsAdmin(ctx) {
		id := ctx.URLParam("id")
		if tool.Isnumber(id) {
			if dbase.DelUser(tool.String2Int(id)) {
				ctx.HTML("<script language='javascript' type='text/javascript'> alert( '删除成功!'); window.location.href='../';</script>")
			} else {
				//EROR
				ctx.HTML("<script language='javascript' type='text/javascript'> alert( '删除失败!'); window.location.href='../';</script>")
			}
		} else {
			pageerror.NotFound(ctx)
		}
	} else {
		pageerror.NotFound(ctx)
	}
}
