package dbase

import (
	"log"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

//Admin Message Page
func ListMessage(gid int) interface{} { //0 Hot else: ALL
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	var resdata /*[30]*/ Discuss
	var result = map[string]map[string]string{}
	if gid != 0 {
		log.Println("后台加载 All Message ....")
		rows, _ := db.Model(&Discuss{}).Rows()
		defer rows.Close()
		i := 0
		for rows.Next() {
			//resdata[i] ==> resdata
			db.ScanRows(rows, &resdata)
			result[strconv.Itoa(i)] = make(map[string]string)
			result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
			result[strconv.Itoa(i)]["uid"] = strconv.Itoa(resdata.UID)
			result[strconv.Itoa(i)]["username"] = resdata.UserName
			result[strconv.Itoa(i)]["group"] = strconv.Itoa(resdata.GroupId)
			result[strconv.Itoa(i)]["posttime"] = resdata.PostTime.String()
			result[strconv.Itoa(i)]["title"] = resdata.Title
			result[strconv.Itoa(i)]["message"] = resdata.Message
			i++
			// do something
		}
	} else {
		//0 热门
		log.Println("后台加载热门....")
		rows, _ := db.Model(Discuss{}).Where("group_id = ?", gid).Rows()

		defer rows.Close()
		i := 0
		for rows.Next() {
			//resdata[i] ==> resdata
			db.ScanRows(rows, &resdata)
			result[strconv.Itoa(i)] = make(map[string]string)
			result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
			result[strconv.Itoa(i)]["uid"] = strconv.Itoa(resdata.UID)
			result[strconv.Itoa(i)]["username"] = resdata.UserName
			result[strconv.Itoa(i)]["group"] = strconv.Itoa(resdata.GroupId)
			result[strconv.Itoa(i)]["posttime"] = resdata.PostTime.String()
			result[strconv.Itoa(i)]["title"] = resdata.Title
			result[strconv.Itoa(i)]["message"] = resdata.Message
			i++
			// do something
		}
	}

	return result
}

func ReplyView(reid int) interface{} {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动迁移
	var resdata /*[30]*/ ReplyMessage
	var result = map[string]map[string]string{}
	rows, _ := db.Model(&ReplyMessage{}).Where("re_id = ?", reid).Rows()
	defer rows.Close()
	i := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &resdata)
		result[strconv.Itoa(i)] = make(map[string]string)
		result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
		result[strconv.Itoa(i)]["message"] = resdata.Message
		result[strconv.Itoa(i)]["ReplyTime"] = resdata.ReplyTime.String()
		result[strconv.Itoa(i)]["username"] = resdata.UserName
		i++
	}
	return result
}

func EditReply(id int, msg string) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&ReplyMessage{}) //自动迁移
	e := db.Model(&ReplyMessage{}).Where(&ReplyMessage{ID: uint(id)}).Update(&ReplyMessage{
		Message: msg,
	}).Error
	if e == nil {
		return true
	}
	return false
}

func AddMeassage(Title, Message string) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移

	if err := db.Create(&Discuss{
		UID:      0,
		GroupId:  1,
		UserName: "Username",
		PostTime: time.Now(),
		Title:    Title,
		Message:  Message,
	}).Error; err != nil {
		//ok
		return false
	}

	return true
}
