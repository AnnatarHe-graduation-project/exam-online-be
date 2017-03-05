package models

import (
	"github.com/jinzhu/gorm"
)

// User is user model
type User struct {
	gorm.Model
	Role      int             `json:"role"` // 10-19 学生 20-29 教师
	Name      string          `json:"name"`
	SchoolID  string          `json:"schoolId"` // 学号，教师号什么的
	Pwd       string          `json:"pwd"`
	Avatar    string          `json:"avatar"`    // 头像
	PaperDone []*StudentPaper `json:"paperDone"` // 完成的卷子，学生
	Papers    []*Paper        `json:"papers"`    // 自己所拥有的卷子, 教师
	News      []*News         `json:"news"`
}
