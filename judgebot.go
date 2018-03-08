package main

import (
	"judgebot/network"
	"os"
	"judgebot/cli"
)

func main() {
	testMode := false
	if len(os.Args) > 1 {
		if os.Args[1] == "--test" || os.Args[1] == "-t" {
			testMode = true
		}
	}

	if testMode {
		cli.Init()
	} else {
		network.InitServer()
	}
}
