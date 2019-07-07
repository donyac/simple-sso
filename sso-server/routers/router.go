package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"simple-sso/sso-server/controllers"
)

/*
sso-server拦截未登录请求
*/
var AccessControlFunc = func(ctx *context.Context) {
	logs.Debug("In access control")
}

func init() {
	//过滤器
	beego.InsertFilter("/*", beego.BeforeExec, AccessControlFunc)
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
}
