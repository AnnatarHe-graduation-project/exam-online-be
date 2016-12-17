package models

import "github.com/jinzhu/gorm"

// News is users put news to this site
type News struct {
	gorm.Model
	Title    string
	Content  string `gorm:"type:text;"`
	Bg       string // 背景大图
	Up       int    `gorm:"default:0;"` // 被赞了多少次
	Down     int    `gorm:"default:0;"` // 被踩了多少次
	UserID   uint
	CourseID []Course `gorm:"many2many:course_has_news;"`
}
