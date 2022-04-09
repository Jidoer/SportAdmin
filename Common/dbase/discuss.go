package dbase

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ListDiscuss(mtype int) interface{} {
	var resdata Discuss
	var result = map[string]map[string]string{}

	if mtype != 0 && mtype != 1 && mtype != 2 {
		return nil
	}
	log.Println(mtype)
	dbtmp, err := gorm.Open("sqlite3", "./data/mydb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动初始化表

	rows, _ := db.Model(&Discuss{}).Where(&Discuss{Group: mtype}).Rows()
	//.Select("id, group_id, uid, answer").
	defer rows.Close()
	i := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &resdata)
		result[strconv.Itoa(i)] = make(map[string]string)
		result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
		result[strconv.Itoa(i)]["uid"] = strconv.Itoa(resdata.UID)
		result[strconv.Itoa(i)]["group"] = strconv.Itoa(resdata.Group)
		result[strconv.Itoa(i)]["posttime"] = resdata.PostTime.String()
		result[strconv.Itoa(i)]["title"] = resdata.Title
		result[strconv.Itoa(i)]["message"] = resdata.Message
		i++
		// do something
	}
	return result
}

func PostMessage(uid, Groupid int, Title, Message string) bool {
	dbtmp, err := gorm.Open("sqlite3", "./data/mydb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动初始化表

	if err := db.Create(&Discuss{
		UID:      uid,
		Group:    Groupid,
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
	dbtmp, err := gorm.Open("sqlite3", "./data/mydb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动初始化表

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
	dbtmp, err := gorm.Open("sqlite3", "./data/mydb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动初始化表.

	if err := db.Create(&ReplyMessage{
		Uid:       uid,
		REID:      reid,
		ReplyTime: time.Now(),
		Message:   message,
	}).Error; err != nil {
		//ok
		return false
	}
	return true
}
