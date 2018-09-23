package controllers

import (
	"log"
)

type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"

	log.Println("------->>", this.Data["path"])
	this.TplName = "home.html"
}

func checkError(err error) {
	log.Println(err)
}
