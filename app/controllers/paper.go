package controllers

import (
	"encoding/json"

	"strconv"

	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// PaperController paper that has many questions
type PaperController struct {
	*revel.Controller
}

// Fetch a paper by paper id
func (p PaperController) Fetch(pid int) revel.Result {
	paper := models.Paper{}
	questions := []models.Question{}
	app.Gorm.Find(&paper, pid)
	app.Gorm.Model(&paper).Association("Questions").Find(&questions)
	paper.Questions = questions

	return p.RenderJson(utils.Response(200, paper, ""))
}

// List 列出所有考卷
func (p PaperController) List() revel.Result {
	papers := []models.Paper{}
	app.Gorm.Find(&papers)
	return p.RenderJson(utils.Response(200, papers, ""))
}

// Add 添加考卷
func (p *PaperController) Add() revel.Result {

	var questions []models.Question
	var courses []models.Course
	title := p.Params.Get("title")
	alert := p.Params.Get("alert")
	questionsID := p.Params.Get("questions")
	score, _ := strconv.Atoi(p.Params.Get("score"))
	coursesStr := p.Params.Get("courses")

	// courses := p.Params.Get("courses")

	var questionsJSON []int
	if err := json.Unmarshal([]byte(questionsID), &questionsJSON); err != nil {
		return p.RenderJson(utils.Response(500, "", err.Error()))
	}
	var coursesJSON []int
	if err := json.Unmarshal([]byte(coursesStr), &coursesJSON); err != nil {
		return p.RenderJson(utils.Response(500, "", err.Error()))
	}

	for _, qid := range questionsJSON {
		question := models.Question{}
		app.Gorm.Find(&question, qid)
		questions = append(questions, question)
	}
	for _, cid := range coursesJSON {
		course := models.Course{}
		app.Gorm.Find(&course, cid)
		courses = append(courses, course)
	}

	hero, err := utils.FileHandler(p.Params.Files["hero"][0])
	if err != nil {
		return p.RenderJson(utils.Response(500, "", err.Error()))
	}

	paper := models.Paper{
		Title:     title,
		Alert:     alert,
		Score:     float32(score),
		Hero:      hero,
		Questions: questions,
		Courses:   courses,
	}

	if err := app.Gorm.Create(&paper).Error; err != nil {
		return p.RenderError(err)
	}

	return p.RenderJson(utils.Response(200, paper, ""))
}

// Avg is get a paper avg score
func (p PaperController) Avg(cid int) revel.Result {
	var count float32
	// var students []models.User
	var studentPaper []models.StudentPaper
	paper := models.Paper{}
	paper.ID = uint(cid)

	app.Gorm.Model(&paper).Select("Score").Find(&studentPaper)

	for _, val := range studentPaper {
		count += val.Score
	}

	avg := count / float32(len(studentPaper))
	return p.RenderJson(utils.Response(200, avg, ""))

}

// Random just for test. it get random question form database
func (p PaperController) Random(cid int) revel.Result {
	// the interfacce returned by api should be equal in front-end

	course := models.Course{}
	var courses []models.Course
	// is there should fetch data by cid with sql coded by myself?
	app.Gorm.Find(&course, cid)
	courses = append(courses, course)

	paper := models.Paper{
		Title:     "test random paper",
		Alert:     "this paper will not get any score",
		Score:     0.00,
		Hero:      "",
		Questions: course.Questions,
		Courses:   courses,
	}

	return p.RenderJson(utils.Response(200, paper, ""))
}
