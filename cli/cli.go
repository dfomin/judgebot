package cli

import (
	"bufio"
	"os"
	"judgebot/commands"
	"strings"
	"fmt"
	"strconv"
)

func Init() {
	const chatMembersCount = 5
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
			result := commands.Judge(params[1:], chatMembersCount)
			fmt.Println(result)
		case "judgeList":
			result := commands.JudgeList(chatMembersCount)
			fmt.Println(result)
		case "judgeVote":
			userID, _ := strconv.Atoi(params[1])
			chatID, _ := strconv.ParseInt(params[2], 10, 64)
			phrase := params[3]
			vote, _ := strconv.ParseBool(params[4])
			fmt.Println(userID, chatID, phrase, vote)
			commands.JudgeVote(userID, chatID, phrase, vote)
		}
	}
}
