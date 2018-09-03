package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	_DB_NAME      = "easyBlog"
	_DB_DRIVER    = "mysql"
	_DB_USERNAME  = "root"
	_DB_PASSWORD  = "xx445602785"
	_DB_ALIASNAME = "default"
)

type Category struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
	Views   int64     `orm:"index"`
	// ArticleTime       time.Time `orm:"index"`
	ArticleCount int64
	// ArticleLastUserId int64
}

type Article struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Create          time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
	CategoryName    string `orm:"index"`
}

/*
后面还可以添加一些username的信息
文章表都是某个username的外键

type User struct{
	Id int64
	Name string
	Passwd string
	ArticleSet int  // user的文章集合
}
*/
func RegisterDB() {
	orm.RegisterModel(new(Category), new(Article))
	// orm.RegisterDriver(_DB_DRIVER, orm.DBMySQL)  // already registed
	orm.RegisterDataBase(_DB_ALIASNAME, _DB_DRIVER,
		_DB_USERNAME+":"+_DB_PASSWORD+"@/"+_DB_NAME+"?charset=utf8")
	orm.Debug = true
	orm.RunSyncdb(_DB_ALIASNAME, false, true)
}

func AddCategoryByName(name string) error {
	o := orm.NewOrm()
	ct := &Category{
		Title:        name,
		Created:      time.Now(),
		ArticleCount: 0,
	}

	err := o.QueryTable("Category").Filter("Title", name).One(ct)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		log.Printf("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {

		// 没有找到记录,执行插入数据操作
		_, ierr := o.Insert(ct)
		if ierr != nil {
			return ierr
		}
	}
	if err == nil {
		// 已经存在记录了则不在进行添加了
		log.Println("this category was existed!")
	}
	return err
}

func AddArticle(name, category, content string) error {
	o := orm.NewOrm()
	newArticle := &Article{
		Title:   name,
		Content: content,
		// Attachment:   "",
		CategoryName: category,
		Create:       time.Now(),
		Updated:      time.Now(),
		ReplyTime:    time.Now(),
		ReplyCount:   0,
	}
	_, err := o.Insert(newArticle)
	return err
}

func UpadateCategoryByName(name string, articleCount int64) error {
	o := orm.NewOrm()
	ct := &Category{
		Title: name,
	}
	// err := GetCategoryByName(name)
	err := o.QueryTable("Category").Filter("Title", name).One(ct)
	if err == orm.ErrNoRows {
		err = AddCategoryByName(name)
		if err != orm.ErrMultiRows {
			log.Println(err)
		}
	}
	if err == nil {
		// 已经存在记录更新记录
		ct.ArticleCount = ct.ArticleCount + articleCount
		if num, err := o.Update(ct); err == nil {
			log.Println("update number:", num)
		}
	}
	return nil
}

func UpdateArticle(title, category, content string, aid int64) error {
	o := orm.NewOrm()
	at := &Article{
		Id: aid,
	}
	err := o.QueryTable("Article").Filter("Id", aid).One(at)
	if err == nil {
		// 已经存在记录更新记录
		at.Title = title
		at.Content = content
		at.Updated = time.Now()
		at.CategoryName = category
		if num, err := o.Update(at); err == nil {
			log.Println("update number:", num)
		}
	}
	return nil
}

func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()
	allData := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&allData)

	return allData, err
}

func GetAllArticles(sort string) ([]*Article, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("article")
	articles := make([]*Article, 0)
	var err error
	switch sort {
	case "DESC":
		_, err = qs.OrderBy("-Create").All(&articles)
	case "ASC":
		_, err = qs.OrderBy("Create").All(&articles)
	default:
		_, err = qs.All(&articles)
	}
	return articles, err
}

func GetAriticleById(id int64) (*Article, error) {
	o := orm.NewOrm()
	article := &Article{Id: id}

	err := o.Read(article)

	if err == orm.ErrNoRows {
		log.Println("查询不到")
	}
	if err == orm.ErrMissPK {
		log.Println("找不到主键")
	}
	return article, nil
}

func DeleteCateGoryById(id int64) error {
	o := orm.NewOrm()
	count, err := o.Delete(&Category{Id: id})
	if err == nil {
		log.Println("被删除的数据：", count)
	}
	return err
}

func DeleteArticleById(id int64) error {
	o := orm.NewOrm()
	count, err := o.Delete(&Article{Id: id})
	if err == nil {
		log.Println("被删除的数据：", count)
	}
	return err
}
