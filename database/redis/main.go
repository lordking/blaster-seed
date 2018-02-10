package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database"
	"github.com/lordking/blaster/database/redis"
	"github.com/lordking/blaster/log"
)

type Reveiver struct{}

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")

	common.InitConfig("redis_exmple", cfgFile)
	log.ReadConfigAt("log")
}

func (d *Reveiver) GetPerson(obj *PersonVO) error {
	log.Debugf("Received a message: %s", common.PrettyObject(obj))

	return nil
}

func main() {

	obj := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	redis := redis.New()
	err := database.Configure("redis", redis)
	common.CheckFatal(err)

	receiver := &Reveiver{}

	p := NewPerson(redis, receiver)

	//设置或新增
	err = p.Set("leking", obj, 1000)
	defer common.CheckFatal(err)
	log.Debugf("set a person: %s", common.PrettyObject(obj))

	//获取
	obj, err = p.Get("leking")
	defer common.CheckFatal(err)
	log.Debugf("get a person: %s", common.PrettyObject(obj))

	// //删除
	err = p.Delete("leking")
	defer common.CheckFatal(err)
	log.Debug("delete a person")

	//订阅
	err = p.Subscribe("person")
	defer common.CheckFatal(err)
	log.Debug("subscribe `person`")

	//发布
	for {
		obj.Phone = fmt.Sprintf("18%d", common.RandInt64(900000000, 999999999))
		log.Debugf("publish a `person`: %s", common.PrettyObject(obj))
		err = p.Publish("person", obj)
		defer common.CheckFatal(err)
		time.Sleep(1e9)
	}

}
