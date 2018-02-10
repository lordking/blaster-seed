package model

import (
	"time"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	Token struct {
		collection *mgo.Collection
	}

	TokenVO struct {
		Id         bson.ObjectId `json:"id" bson:"_id"`
		Token      string        `json:"token" bson:"token"`
		ExpireTime int64         `bson:"expireTime"`
		createTime int64         `bson:"createTime"`
		updateTime int64         `bson:"updateTime"`
	}
)

func (t *Token) Create(obj *TokenVO) error {

	obj.Id = bson.NewObjectId() //生成id
	obj.Token = obj.Id.Hex()
	obj.createTime = time.Now().Unix()
	obj.updateTime = obj.createTime
	obj.ExpireTime = obj.updateTime + 3600

	err := t.collection.Insert(obj)
	if err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}

	return nil
}

func (t *Token) Find(token string) (*TokenVO, error) {

	var result *TokenVO
	objId := bson.ObjectIdHex(token)

	err := t.collection.Find(bson.M{"_id": objId}).One(&result)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	return result, nil
}

func (t *Token) Delete(id string) error {

	objId := bson.ObjectIdHex(id)

	err := t.collection.RemoveId(objId)
	if err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}
	return err
}

/**
 * 清除过期令牌
 * @param  {[type]} t *Token        [description]
 * @return {[type]}   [description]
 */
func (t *Token) ClearExpireTokens() error {

	nowTime := time.Now().Unix()

	_, err := t.collection.RemoveAll(bson.M{"expireTime": bson.M{"$lt": nowTime}})
	if err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}

	return err
}

func NewToken(db *mongo.Mongo) (*Token, error) {

	//获取单例
	err := db.Connect()
	if err != nil {
		err = common.NewError(common.ErrCodeInternal, err.Error())
	}

	collection, err := db.GetCollection("token")
	if err != nil {
		err = common.NewError(common.ErrCodeInternal, err.Error())
	}

	return &Token{collection: collection}, err
}
