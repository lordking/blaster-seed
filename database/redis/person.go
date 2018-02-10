package main

import (
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/redis"
	"github.com/lordking/blaster/log"
)

//Person 用户数据对象
type (
	PersonDelegate interface {
		GetPerson(obj *PersonVO) error
	}

	Person struct {
		delegate PersonDelegate
		db       *redis.Redis
	}

	PersonVO struct {
		Name  string `json:"name" bson:"name"`
		Phone string `json:"phone" bson:"phone"`
	}
)

func (p *Person) Set(key string, obj *PersonVO, expire int) error {

	if err := p.db.Connect(); err != nil {
		return err
	}

	if err := p.db.SetObject(key, obj, expire); err != nil {
		return err
	}

	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Person) Get(key string) (*PersonVO, error) {

	if err := p.db.Connect(); err != nil {
		return nil, err
	}

	obj := new(PersonVO)
	if err := p.db.GetObject(obj, key); err != nil {
		return nil, err
	}

	if err := p.db.Close(); err != nil {
		return nil, err
	}

	return obj, nil
}

func (p *Person) Delete(key string) error {

	if err := p.db.Connect(); err != nil {
		return err
	}

	if err := p.db.DeleteObject(key); err != nil {
		return err
	}

	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Person) Publish(channel string, obj *PersonVO) error {

	if err := p.db.Connect(); err != nil {
		return err
	}

	if err := p.db.PublishObject(channel, obj); err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}

	if err := p.db.Close(); err != nil {
		return err
	}

	return nil
}

func (p *Person) Subscribe(channel string) error {

	if err := p.db.Connect(); err != nil {
		return err
	}

	psc, err := p.db.Subscribe(channel)
	if err != nil {
		return common.NewError(common.ErrCodeInternal, err.Error())
	}

	p.db.Receive(psc)

	go func() {

		for {
			data := <-p.db.ReceiveQueue

			if p.delegate != nil {
				obj := new(PersonVO)
				common.ReadJSON(obj, data)
				if err := p.delegate.GetPerson(obj); err != nil {
					log.Error("Receive error:", err)
				}
			}

		}

	}()

	return nil
}

func NewPerson(db *redis.Redis, delegate PersonDelegate) *Person {
	return &Person{
		db:       db,
		delegate: delegate,
	}
}
