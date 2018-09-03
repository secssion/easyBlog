package controllers

import (
	"easyBlog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	fmt.Println(fmt.Sprint(self.Input()))
	self.Data["PageTitle"] = "Home Page"
	self.Data["IsLogin"] = checkAccount(self.Ctx)

	articleList, err := models.GetAllArticles("DESC")
	checkError(err)

	if len(articleList) > 0 {
		self.Data["at"] = articleList[0]
	} else {
		self.Data["showArticle"] = false
	}

	// self.TplName = "index.tpl"
	self.TplName = "home.html"
}
