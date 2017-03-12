package controllers

import (
	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

type CourseController struct {
	*revel.Controller
}

// List 获取学科列表
func (c CourseController) List() revel.Result {
	courses := []models.Course{}
	papers := []models.Paper{}
	app.Gorm.Find(&courses)

	for index, val := range courses {
		app.Gorm.Model(&val).Related(&papers, "Papers")
		val.Papers = papers
		courses[index] = val
	}

	return c.RenderJson(utils.Response(200, courses, ""))
}
