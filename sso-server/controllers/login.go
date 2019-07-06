package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"simple-sso/sso-server/models"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login/index.tpl"
}

func (c *LoginController) Post() {
	redirectUrl := c.GetString("redirectUrl")
	if len(redirectUrl) == 0 {
		redirectUrl = "http://localhost:8080/index?"
	}
	userName := strings.Trim(c.GetString("username"), " ")
	password := strings.Trim(c.GetString("password"), " ")

	//执行登录行为
	ticket := models.Login(userName, password)

	logs.Debug("ticket=" + ticket)

	if ticket != "" {
		redirectUrl = redirectUrl + "ticket=" + ticket

		logs.Debug(redirectUrl + userName + password)

		c.Redirect(redirectUrl, 302)
	} else {
		c.Redirect("http://localhost:1080/login?redirectURL="+redirectUrl, 302)
	}
}
