package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
	"github.com/lordking/blaster/http"

	"github.com/lordking/blaster-seed/http/simple-mongo/model"
)

type (
	//Person 类声明
	Person struct {
		model *model.Person
	}

	//PersonCreateForm 请求的创建person的json表单
	PersonCreateForm struct {
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone" binding:"required"`
	}

	//PersonUpdateForm 请求的更新person的json表单
	PersonUpdateForm struct {
		Phone string `json:"phone" binding:"required"`
	}
)

//Create 创建用户
func (p *Person) Create(c *gin.Context) {

	var json PersonCreateForm

	err := c.BindJSON(&json)
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	obj := &model.PersonVO{
		Name:  json.Name,
		Phone: json.Phone,
	}

	err = p.model.Create(obj)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, obj)
}

//Find 查询用户
func (p *Person) Find(c *gin.Context) {

	name := c.Param("name")
	result, err := p.model.Find(name)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, result)
}

//Update 更新用户
func (p *Person) Update(c *gin.Context) {

	name := c.Param("name")

	var json PersonUpdateForm

	err := c.BindJSON(&json)
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	obj := &model.PersonVO{
		Phone: json.Phone,
	}

	result, err := p.model.Update(name, obj)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, result)
}

//Delete 删除用户
func (p *Person) Delete(c *gin.Context) {

	name := c.Param("name")

	result, err := p.model.Delete(name)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, result)
}

func NewPerson(db *mongo.Mongo) (*Person, error) {

	model, err := model.NewPerson(db)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	ctrl := &Person{
		model: model,
	}

	return ctrl, nil
}
