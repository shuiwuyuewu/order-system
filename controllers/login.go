package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	token := c.GetSession("user-token")
	if token != nil {
		// check user
		c.TplName = "manage.tpl"
		return
	}

	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	token, _ := c.GetInt("user-token")
	autoLogin := c.GetString("user-auto-login")
	beego.Debug("token: ", token, "auto-login: ", autoLogin)

	if len(autoLogin) != 0 {
		c.SetSession("user-token", token)
	}

	// TODO: check user token
	c.TplName = "manage.tpl"
	return
}
