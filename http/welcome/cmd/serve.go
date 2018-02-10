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
	"github.com/lordking/blaster-seed/http/welcome/api"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/http"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a HTTP server",
	Run: func(cmd *cobra.Command, args []string) {

		config := &http.Config{}
		common.ReadConfigKey("http", config)

		h := http.CreateServer(config)

		// 一个rest服务的简单范例
		welcome := &api.Welcome{}
		group := h.Group("/welcome")
		{
			group.POST("/hello", welcome.Hello)
		}

		h.RunServOnSSL()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
