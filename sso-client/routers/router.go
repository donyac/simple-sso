package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"simple-sso/sso-client/controllers"
	ssoHelper "simple-sso/sso-client/helpers"
)

/*
sso-client拦截未登录请求
*/
var AccessControlFunc = func(ctx *context.Context) {
	logs.Debug("In access control")
	//从cookie里面读取ticket，校验是否存在以及登录
	ticket := ctx.GetCookie("ticket")
	if ticket == "" {
		ticket = ctx.Input.Query("ticket")
	}
	//如果ticket是空，或者未登录
	if (ticket == "") || (!ssoHelper.IsLogin(ticket)) {
		selfUrl := ctx.Request.URL.Path
		loginUrl := "http://localhost:1080/login?redirectURL=http://localhost:8080" + selfUrl
		logs.Debug("ticket=%s loginUrl=%s", ticket, loginUrl)
		ctx.Redirect(302, loginUrl)
	}
	//将ticket设进cookie
	ctx.SetCookie("ticket", ticket)
}

func init() {
	//过滤器
	beego.InsertFilter("/*", beego.BeforeExec, AccessControlFunc)
	beego.Router("/index", &controllers.MainController{}, "get:Get")
	beego.Router("/logout", &controllers.MainController{}, "post:Logout")
	beego.Router("/task/", &controllers.TaskController{}, "get:ListTasks;post:NewTask")
	beego.Router("/task/:id:int", &controllers.TaskController{}, "get:GetTask;put:UpdateTask")
}
