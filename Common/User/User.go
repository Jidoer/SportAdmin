package User

import (
	"github.com/kataras/iris/v12"
	_ "github.com/mattn/go-sqlite3"
	jomesqlite "gormuser/Common/dbase"
	//"JomeSDK/Common/SQL/jomesqlite"
)

/****
* 用户名限制：4~10位大小写字母 或 数字
*
**/

func GetInfoFromPw(user string, pw string) *jomesqlite.UserInfo {
	return jomesqlite.GetInfoFromPw(user, pw)
}
func GetInfoFromCTX(ctx iris.Context) *jomesqlite.UserInfo {
	dump, result := jomesqlite.GetInfoFromCTX(ctx)
	if result != "yes" {
		dump = nil
	}
	return dump
}
func GetUidFromCTX(ctx iris.Context) string {
	return jomesqlite.GetUidFromCTX(ctx)
}
func Reg(user jomesqlite.UserInfo) string {
	return jomesqlite.Reg(user)
}

//鉴权函数
func IsAdmin(ctx iris.Context) bool {
	//can add allow Admin IP
	info, result := jomesqlite.GetInfoFromCTX(ctx)
	if result != "yes" {
		return false
	}
	if info.Vip == 10 {
		return true
	}
	return false
}
