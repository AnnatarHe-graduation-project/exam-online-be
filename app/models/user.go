package models

import (
	"github.com/jinzhu/gorm"
)

// User is user model
type User struct {
	gorm.Model
	Role      int     `json:"role"` // 10-19 学生 20-29 教师
	Name      string  `json:"name"`
	SchoolID  string  `json:"schoolId"` // 学号，教师号什么的
	Pwd       string  `json:"pwd"`
	Avatar    string  `json:"avatar"`                               // 头像
	PaperDone string  `gorm:"type:json" json:"paperDone"`           // 完成的卷子，key是paperID，value是分数
	Papers    []Paper `gorm:"many2many:user_papers;" json:"papers"` // 卷子，什么意思啊
	News      []News  `json:"news"`
}
