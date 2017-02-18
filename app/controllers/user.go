package controllers

import (
	"encoding/json"

	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// UserControl 用户管理
type UserController struct {
	*revel.Controller
}

// Add 添加用户
// http -f POST :9000/auth/register username='AnnatarHe' pwd='aaa' school_id='01111111' role='11'
func (c UserController) Add() revel.Result {
	var username, pwd, schoolID string
	var role int

	c.Params.Bind(&username, "username")
	c.Params.Bind(&pwd, "pwd")
	c.Params.Bind(&schoolID, "school_id")
	c.Params.Bind(&role, "role")

	paperDone, _ := json.Marshal([]map[uint]int{{0: 0}})

	user := models.User{
		Name:      username,
		Pwd:       pwd,
		SchoolID:  schoolID,
		Role:      role,
		PaperDone: string(paperDone),
	}

	if err := app.Gorm.Create(&user).Error; err != nil {
		return c.RenderJson(utils.Response(500, nil, err.Error()))
	}

	return c.RenderJson(utils.Response(200, user, ""))
}

// Login 用户登录 this interface should get username and password for auth
func (c *UserController) Login() revel.Result {
	var username, pwd string
	c.Params.Bind(&username, "username")
	c.Params.Bind(&pwd, "password")
	user := models.User{}

	findUserDb := app.Gorm.Find(&user, map[string]string{
		"Name": username,
		"Pwd":  pwd,
	})

	if err := findUserDb.Error; err != nil {
		return c.RenderError(err)
	}

	if user.Name == "" {
		return c.RenderJson(utils.Response(403, nil, "user should be sign up first"))
	}

	// set session to user

	return c.RenderJson(utils.Response(200, user, ""))
}

// Fetch 获取某个用户数据
func (c UserController) Fetch(uid uint) revel.Result {
	user := models.User{}
	user.ID = uint(uid)
	// app.Gorm.Association("Papers").Find(&user, uid)
	papers := []models.Paper{}
	app.Gorm.Model(&user).Association("Papers").Find(&papers)
	return c.RenderJson(utils.Response(200, user, ""))
}

// FinishedPaper is 完成了某张卷子，记录
func (c UserController) FinishedPaper(pid int) revel.Result {
	// get user id from session

	// and update user paper record
	return c.RenderJson(utils.Response(200, "success", ""))
}
