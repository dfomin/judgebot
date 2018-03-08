package main

import (
	"judgebot/cli"
	"judgebot/network"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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
