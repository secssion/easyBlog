package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) init() {
	this.Data["path"] = this.Ctx.Request.RequestURI
	log.Println(" BaseController PrePare execute------>>>", this.Data["path"])
	// this.Data["Email"] = "445602785@qq.com"
	// this.TplName = "login.html"
}
