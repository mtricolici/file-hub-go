package controllers

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/sessions"
)

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginController struct {
	Ctx iris.Context
}

// GET /login
func (c *LoginController) Get() {
	c.Ctx.View("login.html")
}

// POST /login
func (c *LoginController) Post() {
	username := c.Ctx.FormValue("username")
	password := c.Ctx.FormValue("password")

	if username == "admin" && password == "admin" {
		session := sessions.Get(c.Ctx)
		session.Set("uid", 22)
		c.Ctx.JSON(LoginResponse{Success: true, Message: "ok"})

	} else {
		c.Ctx.JSON(LoginResponse{Success: false, Message: "invalid username or password"})
	}
}
