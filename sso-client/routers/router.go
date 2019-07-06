package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"simple-sso/sso-client/controllers"
)

/*
sso-client拦截未登录请求
 */
var AccessControlFunc = func(ctx *context.Context) {
	logs.Debug("In access control")
	//userName := ctx.Input.Session("userName")
	//if userName == nil {
	//	ctx.Redirect(302, "/login")
	//}
}

func init() {
	//过滤器
	beego.InsertFilter("/*", beego.BeforeExec, AccessControlFunc)

	beego.Router("/", &controllers.MainController{})
}