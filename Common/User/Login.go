package User

import (
	//"encoding/base64"
	"github.com/kataras/iris/v12"
	jomesqlite "gormuser/Common/dbase"	
)

func Login(user string, pw string) (string, jomesqlite.UserInfo) {
	return jomesqlite.Login(user, pw)
}

func IfLogin(ctx iris.Context) bool { //id + key batter than username +key!
	Cookies := ctx.GetCookie("token")
	if Cookies == "" {
		return false
	}
	if jomesqlite.CKToken(Cookies) {
		return true
	}
	return false
}

func EasyLogin(ctx iris.Context) string {
	//兼容
	if IfLogin(ctx) {
		return "yes"
	} else {
		return "no"
	}
}
