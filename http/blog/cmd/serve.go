// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/lordking/blaster-seed/http/blog/api"
	"github.com/lordking/blaster-seed/http/blog/model"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database"
	"github.com/lordking/blaster/database/mongo"
	"github.com/lordking/blaster/http"
	"github.com/lordking/blaster/log"
)

var token *model.Token

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start HTTP serve",
	Run: func(cmd *cobra.Command, args []string) {

		//创建一个数据库访问单例
		mongo := mongo.New()
		err := database.Configure("database", mongo)
		defer common.CheckFatal(err)

		//创建web服务
		config := &http.Config{}
		common.ReadConfigKey("http", config)
		h := http.CreateServer(config)
		h.Router.Use(static.Serve("/", static.LocalFile("./assets", false)))

		user, err := api.NewUser(mongo)
		defer common.CheckFatal(err)
		userGroup := h.Group("/user")
		{
			userGroup.POST("/login", user.Login)
		}

		token, err = model.NewToken(mongo)
		defer common.CheckFatal(err)

		blog, err := api.NewBlog(mongo)
		defer common.CheckFatal(err)
		blogGroup := h.Group("/blog", http.BasicAuth(authorize))
		{
			blogGroup.POST("/new", blog.Create)
			blogGroup.GET("/:start/:limit", blog.Find)
			blogGroup.PUT("/update/:id", blog.Update)
			blogGroup.DELETE("/delete/:id", blog.Delete)
		}

		h.RunServOnSSL()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func authorize(c *gin.Context, authorization string) error {

	username, password, ok := c.Request.BasicAuth()
	if !ok {
		return errors.New("Not found basic authorization!")
	}

	if username != "" {

		result, _ := token.Find(username)
		if result != nil && result.ExpireTime > time.Now().Unix() {
			log.Debug("%s auth ok", username)
			log.Debug("password:%s", password)
			return nil
		}
	}

	str := fmt.Sprintf("%s auth failure", username)
	return errors.New(str)
}
