package routers

import (
	"mytoken/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/data", &controllers.MainController{}, "get:GetData")
	beego.Router("/list", &controllers.MainController{}, "get:GetList")
	beego.Router("/getToken", &controllers.MainController{}, "get:GetToken")
	beego.Router("/getBstk", &controllers.MainController{}, "get:GetBstk")
}
