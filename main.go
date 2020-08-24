package main

import (
	"hcc/clarinet/action/cmd"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/lib/logger"
)

func init() {
	err := logger.Init()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	config.Parser()
	cmd.Init()
}

func main() {
	if cmd.Cmd == nil {
		panic("Init Error!!")
	}

	cmd.Cmd.Execute()
}
