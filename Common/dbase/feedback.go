package dbase

import (
	"gormuser/Common/tool"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

func PostFeedback(email, msg string,bktype int) bool {
	if !tool.IsEmail(email) {
		return false
	}
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&FeedBack{}) //自动迁移
	if err := db.Create(&FeedBack{
		Email: email,
		Msg:   msg,
		BackType: bktype,
	}).Error; err != nil {
		return false
	}

	return true

}
func ListFeedBacks() interface{} {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&FeedBack{}) //自动迁移
	var resdata /*[30]*/ FeedBack
	var result = map[string]map[string]string{}
	rows, _ := db.Model(&FeedBack{}).Rows()
	defer rows.Close()
	i := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &resdata)
		result[strconv.Itoa(i)] = make(map[string]string)

		result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
		result[strconv.Itoa(i)]["email"] = resdata.Email
		result[strconv.Itoa(i)]["message"] = resdata.Msg
		result[strconv.Itoa(i)]["PostTime"] = resdata.PostTime.String()
		result[strconv.Itoa(i)]["type"] = strconv.Itoa(resdata.BackType)
		i++
	}
	return result
}


func EditFeedback(id int, msg string) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&FeedBack{}) //自动迁移
	e := db.Model(&FeedBack{}).Where(&FeedBack{ID: uint(id)}).Update(&FeedBack{
		Msg: msg,
	}).Error
	if e == nil {
		return true
	}
	return false
}
func DelFeedBack(id int) bool {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&FeedBack{}) //自动迁移
	e := db.Model(&FeedBack{}).Where("id = ?", id).Unscoped().Delete(&FeedBack{}).Error
	if e != nil {
		log.Println(strconv.Itoa(id) + ": Del Reply Error!")
		return false
	}
	return true
}
