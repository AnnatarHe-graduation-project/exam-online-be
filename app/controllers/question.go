package controllers

import (
	"strconv"

	"errors"

	"io/ioutil"

	"github.com/AnnatarHe/exam-online-be/app"
	"github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/AnnatarHe/exam-online-be/app/utils"
	"github.com/revel/revel"
	"github.com/tealeg/xlsx"
)

// QuestionController is question controller
type QuestionController struct {
	*revel.Controller
}

// 必须按规定
// | title   | content | answer             					| correct | score | course
// | string  | string  | []{string: string}                     | int     | float | string
// | 谁最帅   | 请问谁最帅 | [{A: 'AnnatarHe'}, {B: 'liang wang'}] | 1		 | 100   | '管理学'
const (
	titleColumn = iota
	contentColumn
	answerColumn
	correctColumn
	scoreColumn
	courseColumn
)

// Add 添加题库，cid是课程ID
func (q *QuestionController) Add(cid int) revel.Result {

	title := q.Params.Get("title")
	content := q.Params.Get("content")
	answers := q.Params.Get("answers")
	correct := q.Params.Get("correct")
	scoreStr := q.Params.Get("score")
	courseStr := q.Params.Get("courses")

	score, _ := strconv.Atoi(scoreStr)
	course, _ := strconv.Atoi(courseStr)

	courses := []models.Course{}
	courseFromDb := models.Course{}
	app.Gorm.Find(&courseFromDb, course)
	courses = append(courses, courseFromDb)

	question := models.Question{
		Title:   title,
		Content: content,
		Answers: answers,
		Correct: correct,
		Score:   score,
		Courses: courses,
	}

	if err := app.Gorm.Create(&question).Error; err != nil {
		return q.RenderJson(utils.Response(500, "", err.Error()))
	}

	return q.RenderJson(utils.Response(200, question, ""))
}

// Fetch find a question by questionID
func (q QuestionController) Fetch(qid int) revel.Result {
	question := models.Question{}
	app.Gorm.Find(&question, qid)
	return q.RenderJson(utils.Response(200, question, ""))
}
func (q QuestionController) FetchAll() revel.Result {
	questions := []models.Question{}
	app.Gorm.Find(&questions)
	return q.RenderJson(utils.Response(200, questions, ""))
}

// AddFromExcel 将文件传入服务中，解析并返回数据
func (q QuestionController) AddFromExcel() revel.Result {

	file, e := q.Params.Files["excel"][0].Open()
	defer file.Close()
	if e != nil {
		revel.INFO.Println(e)
	}
	content, _ := ioutil.ReadAll(file)

	questions, err := decodeExcel(content)

	for _, question := range questions {
		revel.INFO.Println(question)
		if err := app.Gorm.Create(&question).Error; err != nil {
			return q.RenderJson(utils.Response(500, "", err.Error()))
		}
	}

	// err := ioutil.WriteFile("/tmp/gofile.js", content, 0777)
	if err != nil {
		return q.RenderJson(utils.Response(500, "", err.Error()))
	}

	return q.RenderJson(utils.Response(200, "success", ""))
}

// decodeExcel: 解码文件，并存入到数据库
func decodeExcel(buffer []byte) ([]models.Question, error) {
	var questions []models.Question
	xlsxData, err := xlsx.OpenBinary(buffer)
	if err != nil {
		return questions, err
	}

	for _, sheet := range xlsxData.Sheets {
		var question models.Question
		for _, row := range sheet.Rows {
			for index, cell := range row.Cells {
				text, err := cell.String()
				if err != nil {
					return questions, err
				}

				switch index {
				case titleColumn:
					question.Title = text
				case contentColumn:
					question.Content = text
				case answerColumn:
					question.Answers = text
				case correctColumn:
					question.Correct = text
				case scoreColumn:
					i, err := strconv.Atoi(text)
					if err != nil {
						continue
					}
					question.Score = i
				case courseColumn:
					var courses []models.Course
					course := models.Course{Name: text}
					if err := app.Gorm.Find(&course).Error; err != nil {
						question.Courses = courses
					} else {
						question.Courses = append(courses, course)
					}
				default:
					return questions, errors.New("decode error")
				}
			}
			questions = append(questions, question)
		}
	}

	return questions, nil
}
