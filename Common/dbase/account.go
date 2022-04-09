package dbase

import (
	"strconv"

	"github.com/jinzhu/gorm"
)

func ListUser() interface{} {
	dbtmp, err := gorm.Open("sqlite3", "./data/mydb.db")
	if err != nil {
		panic("failed to connect database")
	}
	db = dbtmp
	db.AutoMigrate(&UserInfo{}) //自动初始化表
	var resdata /*[30]*/ UserInfo
	var result = map[string]map[string]string{}
	rows, _ := db.Model(&UserInfo{}).Rows()


	defer rows.Close()

	i := 0
	for rows.Next() {
		//resdata[i] ==> resdata
		db.ScanRows(rows, &resdata)
		result[strconv.Itoa(i)] = make(map[string]string)

		result[strconv.Itoa(i)]["id"] = strconv.Itoa(int(resdata.ID))
		result[strconv.Itoa(i)]["username"] = resdata.Username
		result[strconv.Itoa(i)]["password"] = "*** ***"
		result[strconv.Itoa(i)]["sex"] = strconv.Itoa(resdata.Sex)
		result[strconv.Itoa(i)]["money"] = strconv.Itoa(resdata.Money)
		result[strconv.Itoa(i)]["vip"] = strconv.Itoa(resdata.Vip)
		result[strconv.Itoa(i)]["phone"] = resdata.Phone
		result[strconv.Itoa(i)]["email"] = resdata.Email
		i++
		// do something
	}
	return result
}