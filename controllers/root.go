package controllers

import (
	"github.com/astaxie/beego"
)

type RootController struct {
	beego.Controller
}

func (c *RootController) Get() {
	token := c.GetSession("user-token")
	if token != nil {
		if isAdmin, err := CheckUser(token.(int)); err == nil {
			if isAdmin {
				c.Redirect("/manage", 302)
			} else {
				c.Redirect("/search", 302)
			}
			return
		}
	}

	c.Redirect("/login", 302)
}
