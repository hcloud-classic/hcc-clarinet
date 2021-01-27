package main

import (
	errors "innogrid.com/hcloud-classic/hcc_errors"

	"hcc/clarinet/action/cmd"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/lib/logger"
)

func init() {
	err := logger.Init()
	if err != nil {
		errors.NewHccError(errors.ClarinetInternalInitFail, "logger").Fatal()
	}

	config.Parser()
	cmd.Init()
}

func main() {
	if cmd.Cmd == nil {
		errors.NewHccError(errors.ClarinetInternalInitFail, "cobra").Fatal()
	}

	cmd.Cmd.Execute()
}
