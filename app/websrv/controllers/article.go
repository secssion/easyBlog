package controllers

import (
	"easyBlog/models"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (self *ArticleController) Get() {
	self.Data["PageTitle"] = "article page"
	add := self.Input().Get("add") == "true"
	self.Data["show"] = !add

	aid := self.Input().Get("id")
	id, err := strconv.Atoi(aid)
	checkError(err)

	article, err := models.GetAriticleById(int64(id))
	checkError(err)
	if article != nil {
		self.Data["update"] = true
	}

	op := self.Input().Get("op")
	switch op {
	case "Editor":
		// EditorArticle(article)
		self.Data["show"] = false
	case "Delete":
		DeleteArticleById(int64(id))
	}
	self.Data["at"] = article
	self.Data["showArticle"] = true
	self.TplName = "article.html"
}

func (self *ArticleController) Post() {
	title := self.Input().Get("title")
	category := self.Input().Get("category")
	content := self.Input().Get("content")
	aid := self.Input().Get("id")
	if len(aid) > 0 {
		err := updateCurrentArticle(title, category, content, aid)
		checkError(err)
	} else {
		err := AddNewArticle(title, category, content)
		if err != nil {
			log.Fatal(err)
		}
	}

	articleList, err := models.GetAllArticles("DESC")
	if err != nil {
		log.Println(err)
	}
	self.Redirect(fmt.Sprintf("/article?id=%lld", articleList[0].Id), 301)

	self.TplName = "article.html"
}

func AddNewArticle(title, category, content string) error {
	var err error
	// 如果存在文章分类，则文章的分类类型
	if len(category) > 0 {
		err = models.UpadateCategoryByName(category, 1)
		if err != nil {
			log.Fatal("add category index failed... ", err)
		}
	}

	// 添加文章到数据库之后直接显示
	err = models.AddArticle(title, category, content)
	if err != nil {
		log.Fatal("add category index failed... ", err)
	}
	return err
}

func updateCurrentArticle(title, category, content, aid string) error {
	id, err := strconv.Atoi(aid)
	checkError(err)
	article, err := models.GetAriticleById(int64(id))
	checkError(err)
	if len(article.CategoryName) > 0 && article.CategoryName != category {
		err = models.UpadateCategoryByName(article.CategoryName, -1)
		if err != nil {
			log.Fatal("add category index failed... ", err)
		}
		if len(category) > 0 {
			err = models.UpadateCategoryByName(category, 1)
			if err != nil {
				log.Fatal("add category index failed... ", err)
			}
		}
	}

	err = models.UpdateArticle(title, category, content, int64(id))
	if err != nil {
		log.Fatal("add category index failed... ", err)
	}
	return err
}

func DeleteArticleById(id int64) {
	err := models.DeleteArticleById(id)
	checkError(err)
}
