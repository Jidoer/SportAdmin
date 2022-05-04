package dbase

import (
	"log"

	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" //添加mysql支持
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ListDiscuss(mtype int) interface{} {
	var resdata Discuss
	var result = []map[string]string{}
	var re1 = map[string]string{}

	if mtype != 0 && mtype != 1 && mtype != 2 {
		return nil
	}
	log.Println(mtype)
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	if mtype != 0 {
		//最新
		rows, _ := db.Model(&Discuss{}).Rows()
		//.Select("id, group_id, uid, answer").
		defer rows.Close()
		for rows.Next() {
			db.ScanRows(rows, &resdata)
			re1 = make(map[string]string)
			re1["id"] = strconv.Itoa(int(resdata.ID))
			re1["titletext"] = resdata.UserName
			re1["group"] = strconv.Itoa(resdata.GroupId)
			re1["posttime"] = resdata.PostTime.String()
			re1["title"] = resdata.Title
			re1["message"] = resdata.Message
			result = append(result, re1)
		}
		return result
	} else {
		rows, _ := db.Model(&Discuss{}).Where("group_id = ?", mtype /*Group=0*/).Rows()
		//.Select("id, group_id, uid, answer").
		defer rows.Close()
		i := 0
		for rows.Next() {
			//resdata[i] ==> resdata
			db.ScanRows(rows, &resdata)
			re1 = make(map[string]string)
			re1["id"] = strconv.Itoa(int(resdata.ID))
			re1["titletext"] = resdata.UserName
			re1["group"] = strconv.Itoa(resdata.GroupId)
			re1["posttime"] = resdata.PostTime.String()
			re1["title"] = resdata.Title
			re1["message"] = resdata.Message
			result = append(result, re1)
			i++
			// do something
		}
		return result
	}
}

func PostMessage(uid, Groupid int, Title, Message string) bool {
	uesr, _ := GetInfoFromID(uid)
	var newus UserInfo
	newus = *uesr

	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移

	if err := db.Create(&Discuss{
		UID:      uid,
		GroupId:  Groupid,
		UserName: newus.Username,
		PostTime: time.Now(),
		Title:    Title,
		Message:  Message,
	}).Error; err != nil {
		//ok
		return false
	}

	return true
}

func GetMessageInfo(id int) Discuss {
	var resdata Discuss
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移

	rows, _ := db.Model(&Discuss{}).Where("id=?", id).Rows()
	defer rows.Close()
	i := 0
	for rows.Next() {
		db.ScanRows(rows, &resdata)
		i++
	}
	return resdata
}

func Reply(uid, reid int, message string) bool {
	uesr, _ := GetInfoFromID(uid)
	var newus UserInfo
	newus = *uesr
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动迁移 .

	//x有无msg先不做判断

	if err := db.Create(&ReplyMessage{
		Uid:       uid,
		REID:      reid,
		ReplyTime: time.Now(),
		Message:   message,
		UserName:  newus.Username,
	}).Error; err != nil {
		//ok
		return false
	}
	return true
}

func ListReply(reid int) interface{} {
	var resdata ReplyMessage
	var result = []map[string]string{}
	var re1 = map[string]string{}

	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动迁移

	rows, _ := db.Model(&ReplyMessage{}).Where("re_id=?",reid).Rows()
	//.Select("id, group_id, uid, answer").
	defer rows.Close()
	i := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &resdata)
		re1 = make(map[string]string)
		re1["id"] = strconv.Itoa(int(resdata.ID))
		re1["userId"] = strconv.Itoa(resdata.Uid)
		re1["userName"] = resdata.UserName
		re1["content"] = resdata.Message
		//re1["LikeCount"] = strconv.Itoa(resdata.good)
		//re1["isLike"] = "false"
		//re1["totalCount"] = "0"
		re1["creatTime"] = "刚刚"
		re1["headImg"] = "https://s3.bmp.ovh/imgs/2022/04/13/2157172504dd1793.jpg"

		result = append(result, re1)

		i++
	}

	return result
}

func IsLike(uid, rid int) bool {

	return false
}

func EditMsg(id int, Msg Discuss) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	e := db.Model(&Discuss{}).Where(&Discuss{ID: uint(id)}).Update(&Discuss{
		Title:   Msg.Title,
		Message: Msg.Message,
	}).Error
	if e == nil {
		return true
	}
	return false
}

//ToHot
func ToHot(id int) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	e := db.Model(&Discuss{}).Where(Discuss{ID: uint(id)}).Update("group_id", 0).Error
	if e == nil {
		return true
	}
	return false
}
func NoToHot(id int) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	e := db.Model(&Discuss{}).Where("id=?", id).Update("group_id", 1).Error
	if e == nil {
		return true
	}
	return false
}
func DelReply(id int) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动迁移
	e := db.Model(&ReplyMessage{}).Where("id = ?", id).Unscoped().Delete(&ReplyMessage{}).Error
	if e != nil {
		log.Println(strconv.Itoa(id) + ": Del Reply Error!")
		return false
	}
	return true
}

