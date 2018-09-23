package db

import (
	// "fmt"
	"log"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"time"
)

const (
	IP4        = "127.0.0.1"
	PORT       = "27017"
	USER       = "syeb"
	PWD        = "syeb2014Admin"
	OPERATE_DB = "eBlog"
	TIMEE_OUT  = 60 * time.Second
	POOL_LIMIT = 4096
)

var globalS *mgo.Session

func init() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:     []string{IP4 + ":" + PORT},
		Timeout:   TIMEE_OUT,
		Source:    OPERATE_DB,
		Username:  USER,
		Password:  PWD,
		PoolLimit: POOL_LIMIT,
	})

	if err != nil {
		log.Fatal("create session error", err)
	}

	globalS = session
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	session := globalS.Copy()
	c := session.DB(db).C(collection)
	return session, c
}

func Insert(db, collection string, docs ...interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Insert(docs...)
}

func FindOne(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).All(result)
}

func FindAllSort(db, collection, sort string, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).Sort(sort).All(result)
}

func FindWithPage(db, collection, sort string, page, limit int, query, selector, result interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Find(query).Select(selector).Sort(sort).Skip(page * limit).Limit(limit).All(result)
}

func Update(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	return c.Update(selector, update)
}

func UpdateAll(db, collection string, selector, update interface{}) error {
	session, c := connect(db, collection)
	defer session.Close()
	_, err := c.UpdateAll(selector, update)
	return err
}
