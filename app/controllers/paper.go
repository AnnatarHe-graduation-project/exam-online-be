package controllers

import (
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
	app.Gorm.Find(&paper, pid)

	return p.RenderJson(utils.Response(200, paper, ""))
}

// add a paper just not random
func (p *PaperController) Add() revel.Result {
	var title, alert string
	var score float32
	var questionsID, coursesID []int

	var questions []*models.Question
	var courses []*models.Course
	p.Params.Bind(&title, "title")
	p.Params.Bind(&alert, "alert")
	p.Params.Bind(&score, "score")
	p.Params.Bind(&questionsID, "questions")
	p.Params.Bind(&coursesID, "courses")

	for _, qid := range questions {
		question := models.Question{}
		app.Gorm.Find(&question, qid)
		questions = append(questions, &question)
	}
	for _, cid := range courses {
		course := models.Course{}
		app.Gorm.Find(&course, cid)
		courses = append(courses, &course)
	}

	// there should to deal the pic
	hero := "hero"

	paper := models.Paper{
		Title:     title,
		Alert:     alert,
		Score:     score,
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
	var courses []*models.Course
	// is there should fetch data by cid with sql coded by myself?
	app.Gorm.Find(&course, cid)
	courses = append(courses, &course)

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
