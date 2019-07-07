package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	//初始化一个全局变量对象，存储ticket的缓存，使用beego的memory cache
	ticketCache cache.Cache
	db          *gorm.DB
)

type UserInfo struct {
	Id   uint32
	Name string
}

func init() {
	//每隔 60s 会进行一次过期清理
	//ticketCache, _ = cache.NewCache("memory", `{"interval":60}`)
	ticketCache, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"60"}`)
	var err error
	db, err = gorm.Open("mysql", "root:@(127.0.0.1:3306)/sso?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}

func Login(userName, password string) string {
	//校验用户名密码
	userInfo := User{}
	db.Where("name = ?", userName).First(&userInfo)
	fmt.Println(userInfo)

	if userInfo.Password == password {
		//刷新cache
		ticket := encodeTicket(userName)
		err := ticketCache.Put(ticket, userInfo, 60*5)
		if err != nil {
			return ""
		}
		return ticket
	}
	return ""
}

func encodeTicket(ticket string) string {
	h := md5.New()
	timeStamp := (string)(time.Now().Unix())
	h.Write([]byte(ticket + timeStamp))
	return hex.EncodeToString(h.Sum(nil))
}

/*
获取ticket获取缓存起来的user info
*/
func Get(ticket string) *UserInfo {
	ticket = "ticket-" + ticket
	if ticketCache.IsExist(ticket) {
		userInfo := ticketCache.Get(ticket).(UserInfo)
		return &userInfo
	}
	return nil
}

func Del(ticket string) bool {
	ticket = "ticket-" + ticket
	err := ticketCache.Delete(ticket)
	if err != nil {
		return false
	}
	return true
}
