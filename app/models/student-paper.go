package models

import (
	"github.com/jinzhu/gorm"
)

// StudentPaper is student has done many paper
type StudentPaper struct {
	gorm.Model
	Student        uint    `gorm:"ForeignKey:UserID" json:"studentID"`
	Paper          uint    `gorm:"ForeignKey:PaperID" json:"paperID"`
	PaperContent   Paper   `json:"paper"`
	StudentContent User    `json:"student"`
	Score          float32 `json:"score"`
}
