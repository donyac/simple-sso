package main

import (
	"github.com/astaxie/beego"
	_ "simple-sso/sso-client/routers"
)

func main() {

	beego.Run()
}
