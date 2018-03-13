package cli

import (
	"bufio"
	"fmt"
	"judgebot/commands"
	"os"
	"strconv"
	"strings"
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

		switch strings.ToLower(params[0]) {
		case "judge":
			chatID, _ := strconv.ParseInt(params[1], 10, 64)
			result := commands.Judge(params[2:], chatID, chatMembersCount)
			fmt.Println(result)
		case "judgelist":
			chatID, _ := strconv.ParseInt(params[1], 10, 64)
			result := commands.JudgeList(chatID, chatMembersCount)
			fmt.Println(result)
		case "judgevote":
			userID, _ := strconv.Atoi(params[1])
			chatID, _ := strconv.ParseInt(params[2], 10, 64)
			vote, _ := strconv.ParseBool(params[3])
			phrase := strings.Join(params[4:], " ")
			fmt.Println(userID, chatID, phrase, vote)
			commands.JudgeVote(userID, chatID, phrase, vote)
		}
	}
}
