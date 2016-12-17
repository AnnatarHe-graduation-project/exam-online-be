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
