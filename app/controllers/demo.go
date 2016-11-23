package controllers

import "github.com/revel/revel"

type Demo struct {
	*revel.Controller
}

func (c *Demo) Database() revel.Result {

	return c.RenderJson(map[string]string{"hello": "hello"})
}
