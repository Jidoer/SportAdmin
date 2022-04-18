package dbase

import (
	//"JomeSDK/Common/tool"
	"gormuser/Common/tool/mydes"
	//"database/sql"
	//"encoding/base64"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/*加入去重*/

func Reg(user UserInfo) string {
	if user.Username == "" || user.Password == "" || user.Phone == "" {
		return "no" //禁止出现空Username|password
	}
	var s int
	//dbtmp, err := go rm.Open("sqlite3", "./data/mydb.db")
	dbtmp, err := gorm.Open(dbtype, mydbase)

	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移

	if CheckCF(user.Username) {
		//username = base64.URLEncoding.EncodeToString([]byte(username))
		//email = base64.URLEncoding.EncodeToString([]byte(email))
		lastpass, err := mydes.Encrypt(user.Password, []byte("2fa6c1e9"))
		if err != nil {
			log.Fatal(err)
		}
		if user.Sex == 1 {
			s = 1
		} else {
			s = 0
		}
		if err := db.Create(&UserInfo{
			Username: user.Username,
			Password: lastpass,
			Sex:      s,
			Created:  time.Now(),
			Money:    user.Money,
			Vip:      user.Vip,
			Viptime:  time.Now(),
			Phone:    user.Phone,
			Email:    user.Email,
			Loginip:  user.Loginip,
		}).Error; err != nil {
			//ok
			return "error"
		}
		return "yes"
	} else {

		return "re" //用户名重复
	}

}

/* old
func RegForAddUser(username, password, sex, phone, email, ip string) string {
	if username == "" || password == "" || phone == "" {
		return "no" //禁止出现空Username|password
	}
	var s int
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移

	if CheckCF(username) {
		//username = base64.URLEncoding.EncodeToString([]byte(username))
		//email = base64.URLEncoding.EncodeToString([]byte(email))
		lastpass, err := mydes.Encrypt(password, []byte("2fa6c1e9"))
		if err != nil {
			log.Fatal(err)
		}
		if sex != "0" && sex != "1" {
			if sex == "男" {
				s = 1
			} else {
				s = 0
			}
		} else {
			if sex == "0" {
				s = 0
			} else {
				s = 1
			}
		}
		if err := db.Create(&UserInfo{
			Username: username,
			Password: lastpass,
			Sex:      s,
			Created:  time.Now(),
			Money:    0,
			Vip:      0,
			Viptime:  time.Now(),
			Phone:    phone,
			Email:    email,
			Loginip:  ip,
		}).Error; err != nil {
			//ok
			return "error"
		}
		return "yes"
	} else {

		return "re" //用户名重复
	}

}
*/

func CheckCF(name string) bool {

	//Username

	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动迁移
	println("[ok]")

	rows, _ := db.Model(&UserInfo{}).Where("Username=?", name).Rows()
	defer rows.Close()
	i := 0
	for rows.Next() {
		//db.ScanRows(rows, &resdata)
		i++
	}
	// 现阶段取消对Username的加密!

	if i == 0 {
		return true
	} else {
		return false
	}
}
