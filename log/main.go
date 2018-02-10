package main

import (
	"flag"

	"github.com/lordking/blaster/common"
	"github.com/lordking/blaster/log"
)

var (
	cfgFile string
)

func init() {
	flag.StringVar(&cfgFile, "config", "", "config file.")

	common.InitConfig("mongo_exmple", cfgFile)
	log.ReadConfigAt("log")
}

func main() {

	log.Debug("This is a `test`")
	log.Debugf("This is a '%s'", "test")

	log.Info("This is a `test`")
	log.Infof("This is a '%s'", "test")

	log.Warn("This is a `test`")
	log.Warnf("This is a '%s'", "test")

	err := common.NewError(common.ErrCodeInternal, "a error message!")

	log.Error("This is a error: ", err)
	log.Errorf("This is a error: %s", err)

	log.Fatalf("This is a fatal: %s", err)

	log.Panicf("This is a panic: %s", err)
}
