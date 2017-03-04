package controllers

import (
	"crypto/md5"
	"strconv"

	"encoding/hex"

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

	username := c.Params.Form.Get("username")
	pwd := c.Params.Form.Get("pwd")
	schoolID := c.Params.Form.Get("school_id")
	avatar := c.Params.Form.Get("avatar")
	role, _ := strconv.Atoi(c.Params.Form.Get("role"))

	md5 := md5.New()
	md5.Write([]byte(pwd))

	realPwd := hex.EncodeToString(md5.Sum(nil))

	// paperDone, _ := json.Marshal([]map[uint]int{{0: 0}})

	user := models.User{
		Name:     username,
		Pwd:      realPwd,
		SchoolID: schoolID,
		Avatar:   avatar,
		Role:     role,
	}

	revel.INFO.Println(user)

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
	c.Session["me"] = string(user.ID)

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

// Me get my profile
func (c UserController) Me() revel.Result {
	id, err := strconv.Atoi(c.Session["mee"])
	if err != nil {
		return c.RenderJson(utils.Response(403, "", "login first plz"))
	}
	return c.Fetch(uint(id))
}
