package main

import (
	"database/sql"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mysql"
	"github.com/lordking/blaster/log"
)

type (

	//PersonVO 用户数据对象
	PersonVO struct {
		ID    int
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	//Person person业务模型
	Person struct {
		db *mysql.MySQL
	}
)

//Insert 插入
func (p *Person) Insert(obj *PersonVO) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("INSERT INTO person(name, phone) VALUES(?, ?)")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(obj.Name, obj.Phone)
	lastID, err := result.LastInsertId()
	defer common.CheckFatal(err)

	log.Debugf("Insert result: the `id` of a new row is `%d`", lastID)

}

//FindAll 查询
func (p *Person) FindAll(name string) {

	conn := (p.db.GetConnection()).(*sql.DB)

	var result []PersonVO
	stmt, err := conn.Query("SELECT id, name, phone FROM person WHERE name = ?", name)
	defer stmt.Close()
	defer common.CheckFatal(err)

	for stmt.Next() {
		var obj PersonVO
		err := stmt.Scan(&(obj.ID), &(obj.Name), &(obj.Phone))
		defer common.CheckError(err)

		result = append(result, obj)
	}

	log.Debugf("Find result: %s", common.PrettyObject(result))

}

//UpdateAll 更新
func (p *Person) UpdateAll(name string, obj *PersonVO) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("UPDATE person SET phone=? where name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(obj.Phone, name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debugf("Update result: the sum of effected rows is `%d`", rowsCount)

}

//RemoveAll 删除
func (p *Person) RemoveAll(name string) {

	conn := (p.db.GetConnection()).(*sql.DB)

	stmt, err := conn.Prepare("DELETE FROM person WHERE name=?")
	defer stmt.Close()
	defer common.CheckFatal(err)

	result, err := stmt.Exec(name)
	rowsCount, err := result.RowsAffected()
	defer common.CheckFatal(err)

	log.Debugf("Delete result: the sum of effected rows is `%d`", rowsCount)

}

//NewPerson 实例化
func NewPerson(db *mysql.MySQL) (*Person, error) {

	err := db.Connect()

	return &Person{
		db: db,
	}, err
}
