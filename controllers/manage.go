package controllers

import (
	"github.com/astaxie/beego"
)

type ManageController struct {
	beego.Controller
}

func (c *ManageController) Get() {
	c.TplName = "manage.tpl"
}