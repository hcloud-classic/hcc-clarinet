package main

import (
	"hcc/clarinet/action/cmd"
	"hcc/clarinet/lib/config"
	"hcc/clarinet/lib/errors"
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

	errStack := errors.NewHccErrorStack(errors.NewHccError(errors.ClarinetInternalInitFail, "test1"))
	errStack.Push(errors.NewHccError(errors.ClarinetInternalInitFail, "test2"))
	errStack.Push(errors.NewHccError(errors.ClarinetInternalInitFail, "test3"))

	err := errStack.Pop()

	err.Println()

	errStack.Dump()

	//cmd.Cmd.Execute()
}
