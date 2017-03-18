package controllers

import "github.com/revel/revel"
import "github.com/AnnatarHe/exam-online-be/app/utils"
import "net/http"
import "time"

// Demo just a test
type Demo struct {
	*revel.Controller
}

func (d Demo) Test() revel.Result {

	expires := time.Now().AddDate(0, 0, 1)

	d.SetCookie(&http.Cookie{Name: "foo", Value: "bardsjfjasldk", Expires: expires})

	d.Session["fuck"] = "shit"

	return d.RenderJson(utils.Response(200, "hello", ""))
}
