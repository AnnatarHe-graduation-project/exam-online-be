package models

import (
	"github.com/jinzhu/gorm"
)

/**
 * 所有问题的集合
 */
type Question struct {
	gorm.Model
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Answers  string   `gorm:"type:json;" json:"answers"`
	Correct  string   `gorm:"type:json;" json:"correct"` // 正确答案的index
	HasBug   int      `json:"hasBug"`
	Stared   int      `json:"stared"`
	Score    int      `json:"score"` // 答对这道题，能得的分数
	CourseID []Course `gorm:"many2many:course_has_questions" json:"courses"`
}
