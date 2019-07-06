package main

import (
	_ "simple-sso/sso-server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

