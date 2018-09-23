package models

import (
	"strconv"
	"test"
	"time"

	"easyBlog/models"
)

func TestUserModel_AddUser(t *test.T) {
	user := &UserModel{
		Id_:        strconv.Itoa(time.Location().Second()),
		Name:       "abnerTan",
		Pwd:        "xx445602785",
		Gender:     "male",
		CreateTime: time.Location("Asia/Shanghai"),
		Faico:      "/static/img/tx.png",
		NickName:   "gopher",
		BrithDay:   time.Location(),
		Motto:      "no matter what happend to you ,just go with your dream!",
		Email:      "445602785@163.com",
		Phone:      "13071898332",
		Info:       "if you like, remeber click attention with me!",
	}
	err := user.AddUser(user)
	t.Errorf(" successed execute AddUserFunction")
}
