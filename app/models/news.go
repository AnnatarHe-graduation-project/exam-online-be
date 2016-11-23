package models

import "github.com/jinzhu/gorm"

type News struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text;"`
	Bg      string // 背景大图
	Up      int    `gorm:"default:0;"` // 被赞了多少次
	Down    int    `gorm:"default:0;"` // 被踩了多少次
}
