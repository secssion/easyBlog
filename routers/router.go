package routers

import (
	"easyBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/1", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/editor", &controllers.EditorController{})
}
