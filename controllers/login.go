package controllers

import (
	"github.com/astaxie/beego"
)

type loginFail struct {
	Result       int
	RedirectPath string
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	token, _ := c.GetInt("user-token")
	autoLogin := c.GetString("user-auto-login")
	beego.Debug("token: ", token, "auto-login: ", autoLogin)

	if len(autoLogin) != 0 {
		c.SetSession("user-token", token)
	} else {
		c.DelSession("user-token")
	}

	if isAdmin, err := CheckUser(token); err == nil {
		ret := loginFail{2, "/search"}
		if isAdmin {
			ret.RedirectPath = "/manage"
		}
		c.Data["json"] = &ret
		c.ServeJSON()
		return
	}

	ret := loginFail{1, "/login"}
	c.Data["json"] = &ret
	c.ServeJSON()
	return
}
