package controllers

import (
	"log"
	"strconv"
	"time"

	"easyBlog/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "445602785@qq.com"
	log.Println("------->>", this.Data["path"])
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "445602785@qq.com"
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	log.Println("username:  -->>>", username)
	log.Println("password:  -->>>", password)

	ur := models.NewUserModel()
	user, err := ur.GetBlogByName(username)
	if err != nil {
		log.Println(err)
	}

	if user.Name == username && user.Pwd == password {
		log.Println("user login success!")
	}

	this.TplName = "testAjax.html"
}

func (this *LoginController) findUser(name, pwd string) bool {
	return false
}

func TestUserModel_AddUser() {
	user := &models.UserModel{
		Id_:        strconv.Itoa(time.Now().Second()),
		Name:       "abnerTan",
		Pwd:        "xx445602785",
		Gender:     "male",
		CreateTime: time.Now(),
		Faico:      "/static/img/tx.png",
		NickName:   "gopher",
		BirthDay:   time.Now(),
		Motto:      "no matter what happend to you ,just go with your dream!",
		Email:      "445602785@163.com",
		Phone:      "13071898332",
		Info:       "if you like, remeber click attention with me!",
	}
	err := user.AddUser(user)
	if err != nil {
		log.Println(" failured execute AddUserFunction", err)
	}
	log.Println(" successed execute AddUserFunction")
}
