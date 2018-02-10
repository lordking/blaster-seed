package main

import (
	"flag"

	"github.com/spf13/viper"
	"github.com/lordking/blaster-seed/http/simple-mongo/benchmark"
	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/log"
	"github.com/lordking/blaster/testbox"
)

var (
	methodName string
	perLimit   int
	max        int
	baseURL    string
)

func init() {
	flag.StringVar(&methodName, "m", "", "test case name")
	flag.Parse()

	common.InitConfig("", "./config.yaml")
	log.ReadConfigAt("log")

	perLimit = viper.GetInt("benchmark.perLimit")
	max = viper.GetInt("benchmark.max")
	baseURL = viper.GetString("benchmark.baseURL")
}

func main() {
	limit := 1000
	if max != -1 {
		limit = max
	}
	t := &benchmark.TestCase{
		BaseURL:   baseURL,
		RandLimit: limit,
	}
	testbox.BenchmarkCall(perLimit, max, t, methodName)
}
