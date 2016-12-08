package models

import (
	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title      string
	Alert      string     // 提示信息，警告信息什么的
	Score      float32    // 可获得学分数量
	Hero       string     // 图片
	QuestionID []Question `gorm:"many2many:paper_has_questions;"`
	CourseID   []Course   `gorm:"many2many:course_has_papers;"`
}
