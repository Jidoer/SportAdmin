package dbase

import (
	//"gormuser/Common/User"
	"gormuser/Common/tool/mydes"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

func GetInfoFromPw(user string, pw string) *UserInfo {
	var info UserInfo
	dbtmp, err := gorm.Open("mysql", mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移

	rows, _ := db.Model(&UserInfo{}).Where("username=?", user).Rows()
	defer rows.Close()

	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &info)
		// do something
	}
	//解密👆
	strEncrypted, err := mydes.Decrypt(info.Password, []byte("2fa6c1e9")) //.Encrypt(str, key)
	if strEncrypted == pw {
		info.Password = "Password"
		return &info
	} else {
		return nil
	}

}

func GetInfoFromCTX(ctx iris.Context) (*UserInfo, string) {

	user := ctx.GetCookie("token") //user token

	if user != "" {
		if IfLogin(ctx) {
			ctx.ViewData("iflogin", true) //logined!!!
			////////
			var info UserInfo
			dbtmp, err := gorm.Open("mysql", mydbase)
			if err != nil {
				panic("failed to connect database")
			}
			db = dbtmp
			db.AutoMigrate(&UserInfo{}) //自动迁移

			uid := FromTokenGetID(user)
			if uid == -1 {
				return nil, "no"
			}

			rows, _ := db.Model(&UserInfo{}).Where("id=?", uid).Rows()
			defer rows.Close()
			loop := 0
			for rows.Next() {
				//resdata[i] ==> resdata
				db.ScanRows(rows, &info)
				loop++
				// do something
			}
			if loop == 0 {
				return nil, "nouser"
			}
			/*
				//Uncode UserName
				username, err := base64.URLEncoding.DecodeString(username)
				checkErr(err)
				email, err := base64.URLEncoding.DecodeString(email)
			*/

			if err != nil {
				log.Println(err)
				return nil, err.Error()
			}
			log.Print("得到用户信息: " + string(info.Username))
			//log.Print("UID:" + strconv.Itoa(uid) /*Int2String*/ + "name:" + string(username) + "sex:" + strconv.Itoa(sex) + "created:" + created.String() + "VIP:" + strconv.Itoa(vip) + "Phone:" + strconv.Itoa(phone) + "email:" + string(email) + "loginip:" + loginip)
			return &info, "yes"
		} else {
			//!
			ctx.SetCookieKV("token", "")
		}
	} else {
		//!
		ctx.SetCookieKV("token", "")
		return nil, "no"
	}
	return nil, "no"

}

func GetInfoFromID(uid int) (*UserInfo, string) {
	var info UserInfo
	dbtmp, err := gorm.Open("mysql", mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移

	rows, _ := db.Model(&UserInfo{}).Where("id=?", uid).Rows()
	defer rows.Close()
	loop := 0
	for rows.Next() {
		db.ScanRows(rows, &info)
		loop++
	}
	if loop == 0 {
		return nil, "nouser"
	}
	if err != nil {
		log.Println(err)
		return nil, err.Error()
	}
	return &info, "yes"

}

func GetUidFromCTX(ctx iris.Context) string {
	Info, result := GetInfoFromCTX(ctx)
	if result == "yes" {
		return strconv.Itoa(int(Info.ID))
	} else {
		return "no"
	}
}

func GetMoneyFromCTX(ctx iris.Context) int {
	Info, result := GetInfoFromCTX(ctx)
	if result == "yes" {
		return Info.Money
	} else {
		return 0
	}
}

func Login(user string, pw string) (string, UserInfo) {
	var info UserInfo
	var NULL UserInfo

	dbtmp, err := gorm.Open("mysql", mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移

	rows, _ := db.Model(&UserInfo{}).Where("username=?", user).Rows()
	defer rows.Close()

	loop := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &info)
		loop++
		// do something
	}
	if loop == 0 {
		//无此用户
		log.Println("无此用户!" + user)
		return "nouser", NULL
	}
	//next:
	passwd, err := mydes.Decrypt(info.Password, []byte("2fa6c1e9")) //解密 not.Encrypt(str, key)
	if err != nil {
		log.Println(err)
	}
	if passwd == pw {
		//-----------------------------------------------------------------
		uncodeusername := info.Username //, err := base64.URLEncoding.DecodeString(username)
		if err != nil {
			log.Println(err)
			return "401", NULL
		}
		log.Println("登录成功: " + string(uncodeusername))
		return "yes", info
	} else {
		return "no", NULL

	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteMoney(ctx iris.Context, money int) bool { /*Update The Money For User*/
	//Get UID
	uid := GetUidFromCTX(ctx)
	dbtmp, err := gorm.Open("mysql", mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移
	db.Model(&UserInfo{}).Where("id=?", uid).Update("money", money)
	log.Print("更新余额:" + strconv.Itoa(money))
	return true
}

///--------------------------Tool-----------------------

func DelUser(id int) bool {
	e := db.Model(&UserInfo{}).Where("id = ?", id).Unscoped().Delete(&UserInfo{}).Error
	if e != nil {
		log.Println(strconv.Itoa(id) + ": 验证Error!")
		//func end!
		return false
	}
	return true
}

func AddUser(us UserInfo) bool {
	if Reg(us) == "yes" {
		return true
	}
	return false
}

func EditUser(id int, us UserInfo) bool {
	//e := db.Where("id = ?", id).Update(&UserInfo{}).Error
	dbtmp, err := gorm.Open("mysql", mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移
	e := db.Model(&UserInfo{}).Where("id=?", id).Update(&UserInfo{Username: us.Username, Password: us.Password, Sex: us.Sex, Created: us.Created, Money: us.Money, Vip: us.Vip, Viptime: us.Viptime, Phone: us.Phone, Email: us.Email, Loginip: us.Loginip}).Error
	if e == nil {
		return true
	}
	log.Println(e)
	return false
}
