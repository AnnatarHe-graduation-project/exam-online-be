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

func getRealPwd(pwd string) (realPwd string) {

	md5 := md5.New()
	md5.Write([]byte(pwd))

	realPwd = hex.EncodeToString(md5.Sum(nil))
	return
}

// Add 添加用户
// http -f POST :9000/auth/register username='AnnatarHe' pwd='aaa' school_id='01111111' role='11'
func (c UserController) Add() revel.Result {

	username := c.Params.Form.Get("username")
	pwd := c.Params.Form.Get("pwd")
	schoolID := c.Params.Form.Get("school_id")
	avatar := c.Params.Form.Get("avatar")
	role, _ := strconv.Atoi(c.Params.Form.Get("role"))

	realPwd := getRealPwd(pwd)

	// paperDone, _ := json.Marshal([]map[uint]int{{0: 0}})

	user, err := models.NewUser(app.Gorm, role, username, schoolID, realPwd, avatar)
	if err != nil {
		return c.RenderJson(utils.Response(500, "", err.Error()))
	}

	return c.RenderJson(utils.Response(200, user, ""))
}

// Login 用户登录 this interface should get username and password for auth
func (c UserController) Login() revel.Result {
	username := c.Params.Get("username")
	pwd := c.Params.Get("password")

	realPwd := getRealPwd(pwd)

	user := models.User{}
	findUserDb := app.Gorm.Find(&user, models.User{
		Name: username,
		Pwd:  realPwd,
	})

	if err := findUserDb.Error; err != nil {
		return c.RenderJson(utils.Response(403, "", err.Error()))
	}

	if user.Name == "" {
		return c.RenderJson(utils.Response(403, nil, "user should be sign up first"))
	}

	c.Session["me"] = strconv.Itoa(int(user.ID))
	return c.RenderJson(utils.Response(200, user, ""))
}

// Fetch 获取某个用户数据
func (c UserController) Fetch(uid int) revel.Result {
	user := models.User{}

	if err := app.Gorm.Find(&user, uid).Error; err != nil {
		return c.RenderJson(utils.Response(500, "", "login plz"))
	}

	papers := []models.Paper{}
	paperDone := []models.StudentPaper{}
	news := []models.News{}

	// app.Gorm.Model(&user).Related(&paperDone, "Student")

	if err := app.Gorm.Model(&user).Related(&papers, "Papers").Related(&paperDone, "Student").Related(&news, "News").Error; err != nil {
		return c.RenderJson(utils.Response(500, "", err.Error()))
	}

	for index, val := range paperDone {
		studentPaperDonePapers := models.Paper{}
		app.Gorm.Model(&val).Related(&studentPaperDonePapers, "Paper")
		paperDone[index].PaperContent = studentPaperDonePapers
	}

	// 给老师看的：学生考试成绩统计
	studentPapers := []models.StudentPaper{}
	if len(papers) > 0 {
		for _, p := range papers {
			studentPaperItem := models.StudentPaper{}
			app.Gorm.Model(&p).Related(&studentPapers, "Paper")
			studentPapers = append(studentPapers, studentPaperItem)
		}
	}

	user.PaperDoneByStudent = studentPapers
	user.PaperDone = paperDone
	user.Papers = papers
	user.News = news

	return c.RenderJson(utils.Response(200, user, ""))
}

// FinishedPaper is 完成了某张卷子，记录
func (c UserController) FinishedPaper(pid int) revel.Result {
	uid, _ := strconv.Atoi(c.Session["me"])

	score, _ := strconv.Atoi(c.Params.Get("score"))

	studentPaper := models.StudentPaper{
		Student: uint(uid),
		Paper:   uint(pid),
		Score:   float32(score),
	}

	if err := app.Gorm.Create(&studentPaper).Error; err != nil {
		return c.RenderJson(utils.Response(500, "", err.Error()))
	}

	return c.RenderJson(utils.Response(200, score, ""))
}

// Me get my profile
func (c UserController) Me() revel.Result {
	idStr, _ := c.Session["me"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		revel.INFO.Println(err.Error())
		return c.RenderJson(utils.Response(403, "", "login first plz"))
	}
	return c.Fetch(id)
}
