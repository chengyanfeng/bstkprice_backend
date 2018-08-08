package main

import (
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
	_ "bstkprice_backend/routers"
	"bstkprice_backend/controllers"
	"time"
)

func main() {
	/*logs.SetLogger(logs.AdapterFile, `{"filename":"logs/bskprice_backend.log","daily":false}`)
	logs.Async()
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterConsole, `{"level":0}`)*/
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
