package controllers

import (
	"easyBlog/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	// "github.com/astaxie/beego/context"
)

type CategoryController struct {
	beego.Controller
}

type Cate struct {
	Id           int64
	Title        string
	ArticleCount int64
}

func (self *CategoryController) Get() {
	self.Data["PageTitle"] = "category page"

	operation := self.Input().Get("op")
	switch operation {
	case "AddCategory":
		name := self.Input().Get("category_name")
		AddCategory(name)
	case "DeleteCategory":
		cid := self.Input().Get("id")
		DeleteCategory(cid)
	case "ShowCateGories":
		cid := self.Input().Get("id")
		ShowCateGories(self, cid)
	}
	categoryData, err := models.GetAllCategory()
	if err != nil {
		log.Println(err)
	}
	self.Data["showCategory"] = len(categoryData) > 0
	self.TplName = "category.html"
	self.Data["categories"] = categoryData
}

func AddCategory(name string) {
	if len(name) <= 0 {
		log.Println("请输入分类名称")
	}
	err := models.AddCategoryByName(name)
	if err != nil {
		log.Println(err)
	}
}

func DeleteCategory(id string) {
	Id, err := strconv.Atoi(id)
	checkError(err)
	log.Println(Id)
	err = models.DeleteCateGoryById(int64(Id))
	checkError(err)
}

func ShowCateGories(self *CategoryController, id string) {
	self.Redirect("/article_list?ShowCateGories="+id, 301)
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
	}
}
