package main

import (
	"github.com/lordking/blaster-seed/http/simple-mysql/cmd"
	"github.com/lordking/blaster/common"
)

func main() {
	common.ConfigRuntime()
	cmd.Execute()
}
