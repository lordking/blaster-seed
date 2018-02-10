package main

import (
	"flag"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database"
	"github.com/lordking/blaster/database/mysql"
	"github.com/lordking/blaster/log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")

	common.InitConfig("mysql_exmple", cfgFile)
	log.ReadConfigAt("log")
}

func main() {

	//创建一个数据库访问单例
	mysql := mysql.New()
	err := database.Configure("database", mysql)
	defer common.CheckFatal(err)

	form := &PersonVO{
		Name:  "leking",
		Phone: "18900000000",
	}

	p, err := NewPerson(mysql)
	defer common.CheckFatal(err)

	//插入数据
	p.Insert(form)

	//查询数据
	p.FindAll(form.Name)

	//更新数据
	form.Phone = "13900001111"
	p.UpdateAll(form.Name, form)

	//删除数据
	p.RemoveAll(form.Name)

}
