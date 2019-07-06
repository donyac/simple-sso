package controllers

import (
	"github.com/astaxie/beego"
	"simple-sso/sso-server/models"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Post() {
	ticket := c.Input().Get("ticket")
	models.Del(ticket)
	c.ServeJSON()
}
