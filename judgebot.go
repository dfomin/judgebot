package main

import (
	"bufio"
	"fmt"
	"judgebot/commands"
	"judgebot/network"
	"os"
	"strconv"
	"strings"
)

func InitCLI() {
	reader := bufio.NewReader(os.Stdin)
	for {
		rawCommand, _ := reader.ReadString('\n')
		command := strings.TrimSpace(rawCommand)

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
		case "judgeVote":
			userID, _ := strconv.Atoi(params[1])
			phrase := params[2]
			vote, _ := strconv.ParseBool(params[3])
			fmt.Println(userID, phrase, vote)
			commands.JudgeVote(userID, phrase, vote)
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
