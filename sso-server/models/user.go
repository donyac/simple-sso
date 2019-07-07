package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64"`       // string默认长度64, 使用这种tag重设。
	Password string	`gorm:"size:64"`
	Abstract string
}

// 设置User的表名为`profiles`
func (User) TableName() string {
	return "user"
}
