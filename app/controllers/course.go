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

// Add 添加一个课程
func (c CourseController) Add() revel.Result {
	name := c.Params.Get("name")
	desc := c.Params.Get("desc")

	course := models.Course{
		Name: name,
		Desc: desc,
	}

	if err := app.Gorm.Create(&course).Error; err != nil {
		return c.RenderJson(utils.Response(500, "", err.Error()))
	}
	return c.RenderJson(utils.Response(200, course, ""))
}
