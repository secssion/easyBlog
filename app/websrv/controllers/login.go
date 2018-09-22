package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

const (
	// _MAX_AGE = 1<<32 - 1
	_MAX_AGE = 12
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	fmt.Println("是否进行了 渲染...........")

	isExit := self.Input().Get("isExit") == "true"
	fmt.Println("是否清空 cookies ......", self.Input().Get("isExit"))
	if isExit {
		self.Ctx.SetCookie("username", "", -1, "/")
		self.Ctx.SetCookie("password", "", -1, "/")
		self.Redirect("/", 301)
		return
	}

	self.Data["PageTitle"] = "Admin login"
	self.TplName = "login.html"
}

func (self *LoginController) Post() {

	username := self.Input().Get("username")
	password := self.Input().Get("password")
	autoLogin := self.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = _MAX_AGE
		}

		self.Ctx.SetCookie("username", username, maxAge, "/")
		self.Ctx.SetCookie("password", password, maxAge, "/")
	}
	self.Redirect("/", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	cx, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := cx.Value

	cx, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := cx.Value

	fmt.Println("--->>>username, password", username, password)
	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
