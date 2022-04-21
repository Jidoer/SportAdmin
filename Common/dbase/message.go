package dbase

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

func ListMessage() interface{} {
	dbtmp, err := gorm.Open(dbtype, mydbase)
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&Discuss{}) //自动迁移
	var resdata /*[30]*/ Discuss
	var result = map[string]map[string]string{}
	rows, _ := db.Model(&Discuss{}).Rows()

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
