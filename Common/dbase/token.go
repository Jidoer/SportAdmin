package dbase

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

func SetToken(User UserInfo, token string) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&TokenDB{}) //自动迁移
	dd, _ := time.ParseDuration("240h")
	if IfTokened(int(User.ID)) {
		if err := db.Model(&TokenDB{}).Update(&TokenDB{UserName: User.Username, Token: token, Created: time.Now(), EndTime: time.Now().Add(dd)}).Error; err != nil {
			return false
		}
		return true
	}
	//无已知Token
	if err := db.Create(&TokenDB{UserID: int(User.ID), UserName: User.Username, Token: token, Created: time.Now(), EndTime: time.Now().Add(dd)}).Error; err != nil {
		return false
	}
	return true
}

func IfTokened(uid int) bool {
	var ittoken TokenDB
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&TokenDB{}) //自动迁移
	rows, _ := db.Model(&TokenDB{}).Where("user_id=?", uid).Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &ittoken)
		//有&满足即可
		return true
	}
	return false
}

//验证Token
func CKToken(token string) bool {
	//get token
	var ittoken TokenDB
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&TokenDB{}) //自动迁移
	um := 0
	rows, _ := db.Model(&TokenDB{}).Where("token=?", token).Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &ittoken)
		um++
		//有&满足即可
		return true
	}
	return false
}

func IfLogin(ctx iris.Context) bool {
	Cookies := ctx.GetCookie("token")
	if Cookies == "" {
		return false
	}
	if CKToken(Cookies) {
		return true
	}
	return false
}

func FromTokenGetID(token string) int {
	//get token
	var ittoken TokenDB
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&TokenDB{}) //自动迁移
	um := 0
	rows, _ := db.Model(&TokenDB{}).Where("token=?", token).Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &ittoken)
		um++
		//有&满足即可
		return ittoken.UserID
	}
	return -1
}

/*
//验证Token
func CKToken (user UserInfo,token string) bool{
	//get token
	var ittoken TokenDB
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&TokenDB{}) //自动迁移

	rows, _ := db.Model(&TokenDB{}).Where("userid=?", int(user.ID)).Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &ittoken)
		if ittoken.ID == user.ID{
			if(ittoken.Token==token){
				return true
			}
		}
	}
	return false
}
*/
