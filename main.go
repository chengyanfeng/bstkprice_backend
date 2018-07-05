package main

import (
	_ "bstkprice_backend/routers"
	"github.com/astaxie/beego"
	"bstkprice_backend/controllers"
	"time"
	"github.com/astaxie/beego/context"
)

var ctx *context.Context

func main() {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept")
	beego.BConfig.Listen.HTTPPort = 6001 //端口设置
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
