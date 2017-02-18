package models

import (
	"github.com/jinzhu/gorm"
)

// TeacherPaper is teacher has many paper and paper belongs to teacher
type TeacherPaper struct {
	gorm.Model
	User  *User  `json:"user"`
	Paper *Paper `json:"paper"`
}
