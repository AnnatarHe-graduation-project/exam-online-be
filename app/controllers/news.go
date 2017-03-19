package controllers

import (
	"strconv"

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

	news := []models.News{}
	app.Gorm.Find(&news)
	return n.RenderJson(utils.Response(200, news, ""))
}

// GetOne from models.News
func (n NewsController) GetOne(nid int) revel.Result {
	news := models.News{}
	app.Gorm.Find(&news, nid)
	return n.RenderJson(utils.Response(200, news, ""))
}

// GetTrendings 获取趋势较好的文章
func (n NewsController) GetTrendings() revel.Result {
	return n.GetAll()
}

// Save a news from user request
func (n *NewsController) Save() revel.Result {
	var courses []models.Course

	bgPath, err := utils.FileHandler(n.Params.Files["bg"][0])

	if err != nil {
		return n.RenderJson(utils.Response(500, "", err.Error()))
	}

	title := n.Params.Get("title")
	content := n.Params.Get("content")
	courseStr := n.Params.Get("courses")

	courseID, _ := strconv.Atoi(courseStr)

	c := models.Course{}
	app.Gorm.Find(&c, courseID)
	courses = append(courses, c)

	uid, _ := strconv.Atoi(n.Session["me"])

	news := models.News{
		Title:   title,
		Content: content,
		Bg:      bgPath,
		UserID:  uint(uid),
		Courses: courses,
	}
	if err := app.Gorm.Create(&news).Error; err != nil {
		return n.RenderJson(utils.Response(500, "", err.Error()))
	}

	return n.RenderJson(utils.Response(200, news, ""))

}
