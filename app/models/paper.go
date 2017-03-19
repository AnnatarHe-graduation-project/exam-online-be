package models

import (
	"github.com/jinzhu/gorm"
)

// Paper 考试卷子
type Paper struct {
	gorm.Model
	Title     string         `json:"title"`
	Alert     string         `json:"alert"` // 提示信息，警告信息什么的
	Score     float32        `json:"score"` // 可获得学分数量
	Hero      string         `json:"hero"`  // 图片
	Questions []Question     `gorm:"many2many:paper_has_questions;" json:"questions"`
	Courses   []Course       `gorm:"many2many:course_has_papers;" json:"courses"`
	Teacher   uint           ` json:"teacher"` // 卷子和教师是一对多的关系
	Students  []StudentPaper `json:"students"` // 卷子和学生是多对多的关系
}
