package models

import (
	"github.com/jinzhu/gorm"
)

// Question is all the question
type Question struct {
	gorm.Model
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Answers string    `gorm:"type:json;" json:"answers"`
	Correct string    `gorm:"type:json;" json:"correct"` // 正确答案的index
	HasBug  int       `gorm:"default:'0'" json:"hasBug"`
	Stared  int       `gorm:"default:'0'" json:"stared"`
	Score   int       `gorm:"default:'0'" json:"score"` // 答对这道题，能得的分数
	Courses []*Course `gorm:"many2many:course_has_questions" json:"courses"`
}
