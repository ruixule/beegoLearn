package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "芮旭乐"
	c.Data["Email"] = "1414100586@qq.com"
	c.TplName = "index.tpl"
}
