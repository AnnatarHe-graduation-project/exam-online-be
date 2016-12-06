package models

import (
	"github.com/jinzhu/gorm"
)

// 课程，相当于分类的功能
type Course struct {
	gorm.Model
	Name       string
	Desc       string
	UserID     []User     `gorm:"many2many:course_has_users;"`
	QuestionID []Question `gorm:"many2many:course_has_questions"`
}
