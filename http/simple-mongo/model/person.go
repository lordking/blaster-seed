package model

import (
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
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

//Create 创建用户
func (p *Person) Create(obj *PersonVO) error {

	obj.Id = bson.NewObjectId()

	err := p.collection.Insert(obj)
	if err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}

	return nil
}

//Find 查找用户
func (p *Person) Find(name string) ([]PersonVO, error) {

	var result []PersonVO

	err := p.collection.Find(bson.M{"name": name}).All(&result)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	return result, nil
}

//Update 更新用户
func (p *Person) Update(name string, obj *PersonVO) (*mgo.ChangeInfo, error) {

	//修改
	result, err := p.collection.UpdateAll(bson.M{"name": name}, bson.M{"$set": bson.M{"phone": obj.Phone}})
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	return result, nil
}

//Delete 删除用户
func (p *Person) Delete(name string) (*mgo.ChangeInfo, error) {

	//测试查询
	result, err := p.collection.RemoveAll(bson.M{"name": name})
	if err != nil {
		return nil, common.NewError(common.ErrCodedParams, err.Error())
	}

	return result, nil
}

func NewPerson(db *mongo.Mongo) (*Person, error) {

	//获取单例
	err := db.Connect()
	if err != nil {
		err = common.NewError(common.ErrCodeInternal, err.Error())
	}

	collection, err := db.GetCollection("person")
	if err != nil {
		err = common.NewError(common.ErrCodeInternal, err.Error())
	}

	return &Person{collection: collection}, err
}
