package controllers

import (
	"encoding/json"

	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// UserControl: 用户管理
type UserController struct {
	*revel.Controller
}

// Add: 添加用户
// http -f POST :9000/auth/register username='AnnatarHe' pwd='aaa' school_id='01111111' role='11'
func (c UserController) Add() revel.Result {

	var username, pwd, schoolID string
	var role int

	c.Params.Bind(&username, "username")
	c.Params.Bind(&pwd, "pwd")
	c.Params.Bind(&schoolID, "school_id")
	c.Params.Bind(&role, "role")

	PaperDone, _ := json.Marshal([]map[uint]int{{0: 0}})

	user := models.User{
		Name:      username,
		Pwd:       pwd,
		SchoolID:  schoolID,
		Role:      role,
		PaperDone: string(PaperDone),
	}

	dbUser := app.Gorm.Create(&user)

	if err := dbUser.Error; err != nil {
		return c.RenderJson(utils.Response(500, nil, err.Error()))
	}

	userFromDb := models.User{}

	revel.INFO.Println(user.ID)

	app.Gorm.First(&userFromDb, user.ID)

	return c.RenderJson(utils.Response(200, userFromDb, ""))
}

// Login: 用户登录
func (c UserController) Login() revel.Result {

	var uid int
	c.Params.Bind(&uid, "uid")

	return c.RenderJson(map[string]int{"uid": uid})
}

// Fetch: 获取某个用户数据
func (c UserController) Fetch(uid int) revel.Result {
	revel.INFO.Println(uid)
	u := models.User{}
	user := app.Gorm.Find(&u, uid)
	return c.RenderJson(user)
}

// 完成了某张卷子，记录
func (c UserController) FinishedPaper(pid int) revel.Result {
	return c.RenderJson(map[int]int{})
}
