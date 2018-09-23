package models

import (
	// "time"

	"easyBlog/dao"

	"gopkg.in/mgo.v2/bson"
)

type CategoryModel struct {
	Id_   string `bson:"_id"`
	Name  string `bson:"name"`
	Count int64  `bson:"count"`
}

func NewCategory() *CategoryModel {
	return &CategoryModel{}
}

func (this *CategoryModel) AddUser(cat *CategoryModel) error {
	return db.Insert(DB_NAME, COLLECTION_CATEGORY, cat)
}

func (this *CategoryModel) GetAllBlogs(page int) ([]CategoryModel, error) {
	var CategoryModels []CategoryModel
	err := db.FindAllSort(DB_NAME, COLLECTION_CATEGORY, "-date", nil, nil, &CategoryModels)
	return CategoryModels, err
}

func (this *CategoryModel) GetBlogById(id string) (CategoryModel, error) {
	var CategoryModel CategoryModel
	err := db.FindOne(DB_NAME, COLLECTION_CATEGORY, bson.M{"_id": id}, nil, &CategoryModel)
	return CategoryModel, err
}

func (this *CategoryModel) GetBlogByName(name string) (CategoryModel, error) {
	var CategoryModel CategoryModel
	err := db.FindOne(DB_NAME, COLLECTION_CATEGORY, bson.M{"name": name}, nil, &CategoryModel)
	return CategoryModel, err
}
