package dbase

import (
	"gormuser/Common/tool"

	"github.com/jinzhu/gorm"
)

func PostFeedback(email, msg string) bool {
	if !tool.IsEmail(email){
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
	}).Error; err != nil {
		return false
	}

	return true

}
