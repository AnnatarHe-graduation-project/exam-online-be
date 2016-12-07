package models

import (
	"github.com/jinzhu/gorm"
)

/**
 * 所有问题的集合
 */
type Question struct {
	gorm.Model
	Title    string
	Content  string
	Answers  string `gorm:"type:json;"`
	Right    string `gorm:"type:json;"`
	HasBug   int
	Stared   int
	Score    int      // 答对这道题，能得的分数
	CourseID []Course `gorm:"many2many:course_has_questions"`
}
