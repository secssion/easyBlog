package routers

import (
	"easyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article_list", &controllers.ArticleListController{})
	// beego.Router("/editor", &controllers.CategoryController{})
}
