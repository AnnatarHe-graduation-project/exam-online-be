package controllers

import (
	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
)

// QuestionController: question controller
type QuestionController struct {
	*revel.Controller
}

// Add: 添加题库，cid是课程ID
func (q *QuestionController) Add(cid int) revel.Result {

	var title, content, answer, correct string
	var score int
	var courses []int

	q.Params.Bind(&title, "title")
	q.Params.Bind(&content, "content")
	q.Params.Bind(&answer, "answer")
	q.Params.Bind(&correct, "right")
	q.Params.Bind(&score, "score")
	q.Params.Bind(&courses, "courses")

	coursesFromDB := []models.Course{}

	for _, course := range courses {
		courseFromDB := models.Course{}
		app.Gorm.Find(&courseFromDB, course)
		coursesFromDB = append(coursesFromDB, courseFromDB)
	}

	question := models.Question{
		Title:    title,
		Content:  content,
		Answers:  answer,
		Correct:  correct,
		Score:    score,
		CourseID: coursesFromDB,
	}

	if err := app.Gorm.Create(&question).Error; err != nil {
		return q.RenderJson(utils.Response(500, "", err.Error()))
	}

	return q.RenderJson(utils.Response(200, question, ""))
}

// Fetch: find a question by questionID
func (q QuestionController) Fetch(qid int) revel.Result {
	question := models.Question{}
	app.Gorm.Find(&question, qid)
	return q.RenderJson(utils.Response(200, question, ""))
}
