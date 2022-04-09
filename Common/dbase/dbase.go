package dbase

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var mydbase = "demo:demo@Jidoer..@tcp(hhhh.fun:3306)/demo?parseTime=true"

type UserInfo struct {
	gorm.Model
	//Uid      int //Updata To New SQL: Gorm
	Username string
	Password string
	Sex      int
	Created  time.Time
	Money    int
	Vip      int //如果=0 普通用户 = 1 vip---》viptime可用 =10 --> admin -->time不可用
	Viptime  time.Time
	Phone    string
	Email    string
	Loginip  string
}

type TokenDB struct {
	ID       uint `gorm:"primary_key"`
	UserID   int
	UserName string
	Token    string
	Created  time.Time
	EndTime  time.Time
}

type Discuss struct {
	ID       uint `gorm:"primary_key"`
	UID      int
	Group    int       // 0 Hot 1 New 2 friend
	PostTime time.Time //提交时间
	Title    string
	Message  string
}

type DianZan struct {
	ID        uint `gorm:"primary_key"`
	Uid       int
	DiscussID int
	ZTime     time.Time //点赞时间
}

type ReplyMessage struct {
	ID        uint `gorm:"primary_key"`
	Uid       int
	REID      int //被回复消息ID
	ReplyTime time.Time
	Message   string
}
