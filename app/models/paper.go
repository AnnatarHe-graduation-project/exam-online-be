package models

import (
	"github.com/jinzhu/gorm"
)

type Paper struct {
	gorm.Model
	Title      string
	Alert      string     // 提示信息，警告信息什么的
	Score      float32    // 可获得学分数量
	QuestionID []Question `gorm:"many2many:paper_has_question;"`
}
