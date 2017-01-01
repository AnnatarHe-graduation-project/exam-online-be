package controllers

import (
	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// NewsController is news controller 😂
type NewsController struct {
	*revel.Controller
}

// GetAll from models.News
func (n NewsController) GetAll() revel.Result {
	news := app.Gorm.Find(&models.News{})
	return n.RenderJson(news)
}

// GetOne from models.News
func (n NewsController) GetOne(nid int) revel.Result {

	news := app.Gorm.Find(&models.News{}, nid)
	return n.RenderJson(utils.Response(200, news, ""))

}

// Save a news from user request
func (n *NewsController) Save(uid int) revel.Result {
	var title, content string
	var coursesID []int
	var courses []models.Course
	// 还有个Bg，背景大图不知道怎么弄

	n.Params.Bind(&title, "title")
	n.Params.Bind(&content, "content")
	n.Params.Bind(&coursesID, "courses")

	for i := 0; i < len(coursesID); i++ {
		c := app.Gorm.Find(&models.Course{}, coursesID[i])
		append(courses, c)
	}

	news := models.News{
		Title:    title,
		Content:  content,
		Bg:       "nil",
		UserID:   uint(uid),
		CourseID: courses,
	}
	if err := app.Gorm.Create(&news).Error; err != nil {
		return n.RenderJson(utils.Response(500, "", err.Error()))
	}

	return n.RenderJson(utils.Response(200, news, ""))

}
