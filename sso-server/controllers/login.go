package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login/index.tpl"
	//c.Ctx.WriteString("hello")
}

type person struct {
	Name string
	Age  int
}

func (c *LoginController) Post() {
	//c.Ctx.WriteString("hello")
	//c.TplName = "login/index.tpl"
	mystruct := person{Name: "dongnan", Age: 10}
	c.Data["json"] = &mystruct
	c.ServeJSON()
}
