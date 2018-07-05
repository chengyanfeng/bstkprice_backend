package main

import (
	_ "bstkprice_backend/routers"
	"github.com/astaxie/beego"
	"bstkprice_backend/controllers"
	"time"
	"github.com/astaxie/beego/plugins/cors"
)



func main() {


	beego.BConfig.Listen.HTTPPort = 6001 //端口设置
	beego.BConfig.RecoverPanic = true
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
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
