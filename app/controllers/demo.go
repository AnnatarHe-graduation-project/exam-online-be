package controllers

import (
	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/revel/revel"
)

// Demo just a test
type Demo struct {
	*revel.Controller
}

// Database provider a method to test database is connected or not
func (c *Demo) Database() revel.Result {

	user := models.User{
		Role:     1,
		Name:     "demo user111",
		SchoolID: "03313138",
		Pwd:      "password",
		Avatar:   "a path",
		News:     []models.News{{Title: "just a test1111", Content: "hello world", Bg: "/some/path/here"}},
	}
	if err := app.Gorm.Create(&user).Error; err != nil {
		revel.INFO.Fatalln(err)
		return c.RenderJson(err)
	}

	return c.RenderJson(app.Gorm.First(&user))

}
