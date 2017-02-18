package models

import (
	"github.com/jinzhu/gorm"
)

// Course 课程，相当于分类的功能
type Course struct {
	gorm.Model
	Name      string     `json:"name"`
	Desc      string     `json:"desc"`
	News      []News     `gorm:"many2many:course_has_news;" json:"news"`
	Papers    []Paper    `gorm:"many2many:course_has_papers;" json:"papers"`
	Questions []Question `gorm:"many2many:course_has_questions;" json:"questions"`
}
