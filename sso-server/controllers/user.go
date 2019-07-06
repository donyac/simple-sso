package controllers

import (
	"github.com/astaxie/beego"
	"simple-sso/sso-server/models"
)

type UserController struct {
	beego.Controller
}

type Result struct {
	Code int
	Msg  string
}

func (c *UserController) Get() {
	//appKey := c.Input().Get("appKey")
	ticket := c.Input().Get("ticket")
	userPtr := models.Get(ticket)

	if userPtr != nil {
		c.Data["json"] = *userPtr
	}

	//todo 先使用假数据，稍后从数据库读取数据
	c.Data["json"] = struct {
		Name string
		Age int
	}{
		"donyac",
		11,
	}
	c.ServeJSON()
}
