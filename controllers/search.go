package controllers

import (
	"github.com/astaxie/beego"
)

// SearchController operations for Search
type SearchController struct {
	beego.Controller
}

func (c *SearchController) Get() {
	c.TplName = "search.tpl"
}
