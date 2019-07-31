package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	logs.Info("loan-picture-web main  ")
	c.TplName = "index.tpl"
}
