package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lordking/blaster-seed/http/blog/model"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
	"github.com/lordking/blaster/http"
)

type (

	//Blog controller声明
	Blog struct {
		model *model.Blog
	}

	//BlogCreateForm 创建日志的json协议
	BlogCreateForm struct {
		Subject string `json:"subject" binding:"required"`
		Blog    string `json:"blog" binding:"required"`
		Author  string `json:"author" binding:"required"`
	}

	//BlogUpdateForm 更新日志的json协议
	BlogUpdateForm struct {
		Subject string `json:"subject" `
		Blog    string `json:"blog" `
	}
)

//Create 创建日志
func (b *Blog) Create(c *gin.Context) {

	var form BlogCreateForm

	err := c.BindJSON(&form)
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	obj := &model.BlogVO{Subject: form.Subject, Blog: form.Blog, Author: form.Author}
	err = b.model.Create(obj)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, obj)
}

//Find 查找日志
func (b *Blog) Find(c *gin.Context) {

	start, err := strconv.Atoi(c.Param("start"))
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	limit, err := strconv.Atoi(c.Param("limit"))
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	result, err := b.model.Find(start, limit)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, result)
}

//Update 更新日志
func (b *Blog) Update(c *gin.Context) {

	id := c.Param("id")
	var form BlogUpdateForm

	err := c.BindJSON(&form)
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	obj := &model.BlogVO{Subject: form.Subject, Blog: form.Blog}

	err = b.model.Update(id, obj)
	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, "ok")
}

//Delete 删除日志
func (b *Blog) Delete(c *gin.Context) {

	id := c.Param("id")
	err := b.model.Delete(id)

	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
		return
	}

	http.JSONResponse(c, 200, "ok")
}

func NewBlog(mongo *mongo.Mongo) (*Blog, error) {

	model, err := model.NewBlog(mongo)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	ctrl := &Blog{
		model: model,
	}

	return ctrl, nil
}
