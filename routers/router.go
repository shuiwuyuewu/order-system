package routers

import (
	"order-system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
}
