package main

import (
	clarinetInit "hcc/clarinet/init"
)

func main() {
	err := clarinetInit.MainInit()
	if err != nil {
		panic(err)
	}
}