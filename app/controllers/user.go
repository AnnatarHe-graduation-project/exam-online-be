package controllers

import (
	"github.com/revel/revel"
)

/**
 * User这个Controller上的
 */
type User struct {
	*revel.Controller
}

/**
 * 添加一个用户的方法
 */
func (c User) Add() revel.Result {

	return c.RenderJson(map[string]string{"hello": "hello"})
}

/**
 * 用户登陆
 */
func (c User) Login() revel.Result {

	// var uid, role, school_id, News_id int
	// var name, pwd, avatar string
	var uid int
	c.Params.Bind(&uid, "uid")

	return c.RenderJson(map[string]int{"uid": uid})
}
