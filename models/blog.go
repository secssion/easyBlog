package models

import (
	"time"

	"easyBlog/dao"

	"gopkg.in/mgo.v2/bson"
)

type BlogModel struct {
	Id_         string         `bson:"_id"`
	UserId      string         `bson:"user_id"`
	CategoryId  []string       `bson:"category_id"`
	Name        string         `bson:"name"`
	PublishTime time.Time      `bson:"publish_time"`
	UpdateTime  time.Time      `bson:"update_time"`
	Origin      string         `bson:"orgin"` // origin content
	Parse       string         `bson:"parse"` // parse content
	Comments    []CommentModel `bson:"comments"`
	Tags        []TagModel     `bson:"tags"`
	BrowseCount int64          `bson:"browse_count"`
	PraiseCount int64          `bson:"praise_count"`
}

type CommentModel struct {
	Content string `bson:"yk_comment"`
	Eamil   string `bson:"yk_email"`
	Name    string `bson:"yk_name"`
	Gender  int64  `bson:"yk_gender"`
}

type TagModel struct {
	TagName string `bson:"tagName"`
	TagId   string `bson:"tagId"`
}

func NewBlog() *BlogModel {
	return &BlogModel{}
}

func (this *BlogModel) AddUser(blog *BlogModel) error {
	return db.Insert(DB_NAME, COLLECTION_BLOG, blog)
}

func (this *BlogModel) GetAllBlogs(page int) ([]BlogModel, error) {
	var blogs []BlogModel
	err := db.FindAllSort(DB_NAME, COLLECTION_BLOG, "-date", nil, nil, &blogs)
	return blogs, err
}

func (this *BlogModel) GetBlogById(id string) (BlogModel, error) {
	var blog BlogModel
	err := db.FindOne(DB_NAME, COLLECTION_BLOG, bson.M{"_id": id}, nil, &blog)
	return blog, err
}

func (this *BlogModel) GetBlogByName(name string) (BlogModel, error) {
	var blog BlogModel
	err := db.FindOne(DB_NAME, COLLECTION_BLOG, bson.M{"name": name}, nil, &blog)
	return blog, err
}
