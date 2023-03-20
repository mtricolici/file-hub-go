package controllers

import (
	"github.com/kataras/iris/v12"

	"github.com/kataras/iris/v12/sessions"
)

type HomeController struct {
	Ctx iris.Context
}

// GET /
func (c *HomeController) Get() {
	user_id, _ := sessions.Get(c.Ctx).GetInt("uid")
	if user_id < 1 {
		c.Ctx.Redirect("/login")
	} else {
		c.Ctx.ViewData("Username", "admin")
		c.Ctx.ViewData("Id", user_id)
		c.Ctx.View("home.html")
	}
}
