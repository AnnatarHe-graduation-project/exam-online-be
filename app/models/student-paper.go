package models

import (
	"github.com/jinzhu/gorm"
)

// StudentPaper is student has done many paper
type StudentPaper struct {
	gorm.Model
	Student *User   `json:"student"`
	Paper   *Paper  `json:"paper"`
	Score   float32 `json:"score"`
}
