package controllers

import (
	"log"
)

type EditorController struct {
	BaseController
}

func (c *EditorController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	log.Println("enter log html")
	c.TplName = "editor.html"
}
