package controllers

import (
	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// UserControl: 用户管理
type UserControl struct {
	*revel.Controller
}

// Add: 添加用户
func (c UserControl) Add() revel.Result {

	var username, pwd, schoolID string
	var role int

	c.Params.Bind(&username, "username")
	c.Params.Bind(&pwd, "pwd")
	c.Params.Bind(&schoolID, "school_id")
	c.Params.Bind(&role, "role")

	user := models.User{
		Name:     username,
		Pwd:      pwd,
		SchoolID: schoolID,
		Role:     role,
	}

	if err := app.Gorm.Create(&user).Error; err != nil {
		return c.RenderJson(utils.Response(500, nil, err.Error()))
	}

	return c.RenderJson(utils.Response(200, app.Gorm.First(&models.User{}, app.Gorm.RowsAffected), ""))
}

// Login: 用户登录
func (c UserControl) Login() revel.Result {

	var uid int
	c.Params.Bind(&uid, "uid")

	return c.RenderJson(map[string]int{"uid": uid})
}

// Fetch: 获取某个用户数据
func (c UserControl) Fetch(uid int) revel.Result {
	revel.INFO.Println(uid)
	u := models.User{}
	user := app.Gorm.Find(&u, uid)
	return c.RenderJson(user)
}

// 完成了某张卷子，记录
func (c UserControl) FinishedPaper(pid int) revel.Result {
	return c.RenderJson(map[int]int{})
}
