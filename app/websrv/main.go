package main

import (
	"easyBlog/models"
	_ "easyBlog/app/websrv/routers"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
