package routers

import (
	"bstkprice_backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/data", &controllers.MainController{}, "get:GetData")

}
