package main

import (
	"bufio"
	"fmt"
	"judgebot/commands"
	"judgebot/network"
	"os"
	"strings"
)

func InitCLI() {
	reader := bufio.NewReader(os.Stdin)
	for {
		command, _ := reader.ReadString('\n')
		params := strings.Split(command, " ")

		if len(params) == 0 {
			continue
		}

		switch params[0] {
		case "judge":
			result := commands.Judge(params[1:])
			fmt.Println(result)
		case "judgeList":
			result := commands.JudgeList()
			fmt.Println(result)
		}
	}
}

func main() {
	testMode := false
	if len(os.Args) > 1 {
		if os.Args[1] == "--test" || os.Args[1] == "-t" {
			testMode = true
		}
	}

	if testMode {
		InitCLI()
	} else {
		network.InitServer()
	}
}
