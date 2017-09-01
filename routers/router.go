package routers

import (
	"order-system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.RootController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/search", &controllers.SearchController{})
	beego.Router("/manage", &controllers.ManageController{})
}
