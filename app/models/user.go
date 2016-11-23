package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Role     int
	Name     string
	SchoolID string // 学号，教师号什么的
	Pwd      string
	Avatar   string  // 头像
	Papers   []Paper `gorm:"many2many:user_papers;"`
	News     []News
}
