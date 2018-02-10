package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database/mongo"
	"github.com/lordking/blaster/http"

	"github.com/lordking/blaster-seed/http/blog/model"
)

type (
	//User controller声明
	User struct {
		token *model.Token
	}

	//UserLoginForm 登录的json协议
	UserLoginForm struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

//Login 用户登录
func (u *User) Login(c *gin.Context) {

	var json UserLoginForm

	err := c.BindJSON(&json)
	if err != nil {
		http.JSONResponse(c, 403, err)
		return
	}

	if json.Username != "admin" && json.Password != "admin" {
		http.JSONResponse(c, 401, "用户名或密码错误")
		return
	}

	obj := &model.TokenVO{}
	err = u.token.Create(obj)
	if err != nil {
		err := err.(*common.Error)
		http.JSONResponse(c, err.Code, err.Message)
	}

	u.token.ClearExpireTokens() //清除已退休的token
	http.JSONResponse(c, 200, obj)
}

func NewUser(mongo *mongo.Mongo) (*User, error) {

	token, err := model.NewToken(mongo)
	if err != nil {
		return nil, common.NewError(common.ErrCodeInternal, err.Error())
	}

	ctrl := &User{
		token: token,
	}

	return ctrl, nil
}
