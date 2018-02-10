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
	"github.com/spf13/cobra"
	"github.com/lordking/blaster-seed/http/simple-mysql/api"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/database"
	"github.com/lordking/blaster/database/mysql"
	"github.com/lordking/blaster/http"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start HTTP serve",
	Run: func(cmd *cobra.Command, args []string) {

		//创建一个数据库访问单例
		mysql := mysql.New()
		err := database.Configure("database", mysql)
		defer common.CheckFatal(err)

		//创建web服务
		config := &http.Config{}
		common.ReadConfigKey("http", config)
		h := http.CreateServer(config)

		person, err := api.NewPerson(mysql)
		defer common.CheckError(err)

		group := h.Group("/person")
		{
			group.POST("/new", person.Create)
			group.GET("/:name", person.Find)
			group.PUT("/update/:name", person.Update)
			group.DELETE("/delete/:name", person.Delete)
		}

		h.RunServOnSSL()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
