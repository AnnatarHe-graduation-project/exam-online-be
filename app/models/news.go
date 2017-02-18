package models

import "github.com/jinzhu/gorm"

// News is users put news to this site
type News struct {
	gorm.Model
	Title   string   `json:"title"`
	Content string   `gorm:"type:text;" json:"content"`
	Bg      string   `json:"bg"`                     // 背景大图
	Up      int      `gorm:"default:0;" json:"up"`   // 被赞了多少次
	Down    int      `gorm:"default:0;" json:"down"` // 被踩了多少次
	User    User     `json:"user"`
	Courses []Course `gorm:"many2many:course_has_news;"`
}
