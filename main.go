package main

import (
	_ "mytoken/routers"
	"github.com/astaxie/beego"
	"mytoken/controllers"
	"time"
)

func main() {

	beego.BConfig.Listen.HTTPPort = 6001                     //端口设置
	beego.BConfig.RecoverPanic = true
	RunTime := controllers.MainController{}
	go TimeGetToken(RunTime)
	beego.Run()
}

func TimeGetToken(c controllers.MainController) {
	ticker := time.NewTicker(60 * time.Second)
	for _ = range ticker.C {
		c.GetToken()
		c.GetBstk()
	}
}
