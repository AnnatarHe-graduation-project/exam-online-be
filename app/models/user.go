package models

import (
	"github.com/jinzhu/gorm"
)

// User is user model
type User struct {
	gorm.Model
	Role      int
	Name      string
	SchoolID  string // 学号，教师号什么的
	Pwd       string
	Avatar    string  // 头像
	PaperDone string  `gorm:"type:json"`              // 完成的卷子，key是paperID，value是分数
	PaperID   []Paper `gorm:"many2many:user_papers;"` // 卷子，什么意思啊
	NewsID    []News
}
