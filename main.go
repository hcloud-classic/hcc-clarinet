package main

import (
	Clarinet "hcc/clarinet/cmd"
	"hcc/clarinet/lib/config"
)

func init() {
	config.Parser()
	Clarinet.Init()
}

func main() {
	if Clarinet.Cmd == nil {
		panic("Init Error!!")
	}

	Clarinet.Cmd.Execute()
}
