package controllers

import (
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Add() revel.Result {

	return c.RenderJson(map[string]string{"hello": "hello"})
}

func (c User) Login() revel.Result {

	var uid int
	c.Params.Bind(&uid, "uid")

	return c.RenderJson(map[string]int{"uid": uid})
}
