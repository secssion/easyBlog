package controllers

import (
	"easyBlog/models"
	"github.com/astaxie/beego"
	"log"
)

type ArticleListController struct {
	beego.Controller
}

func (self *ArticleListController) Get() {
	self.Data["PageTitle"] = "article list page"
	articleList, err := models.GetAllArticles("DESC")
	if err != nil {
		log.Println(err)
	}
	self.Data["IsShowArticles"] = len(articleList) > 0
	self.Data["articles"] = articleList
	self.TplName = "showArticles.html"
}
