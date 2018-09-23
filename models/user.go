package models

import (
	"time"

	"easyBlog/dao"

	"gopkg.in/mgo.v2/bson"
)

type UserModel struct {
	Id_        string    `bson:"_id"`
	Name       string    `bson:"name"`
	Pwd        string    `bson:"pwd`
	Gender     string    `bson:"gender"`
	CreateTime time.Time `bson:"create_time"`
	BirthDay   time.Time `bson:"birth_day"`
	Faico      string    `bson:"facio"`
	NickName   string    `bson:"nickname"`
	Motto      string    `bson:"motto"`
	Email      string    `bson:"eamil"`
	Phone      string    `bson:"phone"`
	Info       string    `bson:"info"`
}

func NewUserModel() *UserModel {
	return &UserModel{}
}

func (this *UserModel) AddUser(user *UserModel) error {
	return db.Insert(DB_NAME, COLLECTION_USER, user)
}

func (this *UserModel) GetAllBlogs(page int) ([]UserModel, error) {
	var users []UserModel
	err := db.FindAllSort(DB_NAME, COLLECTION_USER, "-date", nil, nil, &users)
	return users, err
}

func (this *UserModel) GetBlogById(id string) (UserModel, error) {
	var user UserModel
	err := db.FindOne(DB_NAME, COLLECTION_USER, bson.M{"_id": id}, nil, &user)
	return user, err
}

func (this *UserModel) GetBlogByName(name string) (UserModel, error) {
	var user UserModel
	err := db.FindOne(DB_NAME, COLLECTION_USER, bson.M{"name": name}, nil, &user)
	return user, err
}
