package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	ssoHelper "simple-sso/sso-client/helpers"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	ticket := c.Ctx.GetCookie("ticket")
	userInfoPtr := ssoHelper.GetUserInfo(ticket)
	fmt.Println(*userInfoPtr)
	if userInfoPtr != nil {
		c.Data["UserName"] = userInfoPtr.Name
	}
	c.TplName = "index.html"
	c.Render()
}

func (c *MainController) Logout() {
	ticket := c.Ctx.GetCookie("ticket")
	ssoHelper.Logout(ticket)
}
