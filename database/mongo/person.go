package main

import (
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
	"github.com/lordking/blaster/log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Person struct {
		collection *mgo.Collection
	}

	//Person 用户数据对象
	PersonVO struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Phone string        `json:"phone" bson:"phone"`
	}
)

func (p *Person) insert(obj *PersonVO) (err error) {

	obj.Id = bson.NewObjectId()
	err = p.collection.Insert(obj)

	log.Debugf("Insert result: %s", common.PrettyObject(obj))
	return
}

func (p *Person) findAll(name string) (result []PersonVO, err error) {

	err = p.collection.Find(bson.M{"name": name}).All(&result)

	log.Debugf("Find result: %s", common.PrettyObject(result))
	return
}

func (p *Person) updateAll(name string, obj *PersonVO) (result *mgo.ChangeInfo, err error) {

	result, err = p.collection.UpdateAll(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": obj.Phone}})

	log.Debugf("Update result: %s", common.PrettyObject(result))
	return
}

func (p *Person) removeAll(name string) (result *mgo.ChangeInfo, err error) {

	result, err = p.collection.RemoveAll(bson.M{"name": name})

	log.Debugf("Remove result: %s", common.PrettyObject(result))
	return
}

func NewPerson(db *mongo.Mongo) (p *Person, err error) {

	if err = db.Connect(); err != nil {
		return
	}

	collection, err := db.GetCollection("person")
	if err != nil {
		return
	}

	p = new(Person)
	p.collection = collection

	return
}
