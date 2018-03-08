package network

import (
	"judgebot/commands"
	"judgebot/private"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
	"judgebot/database"
)

func InitServer() {
	bot, err := tgbotapi.NewBotAPI(private.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://pigowl.com:8443/"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServeTLS(":8443", "fullchain.pem", "privkey.pem", nil)

	for update := range updates {
		chatMembersCount, err := bot.GetChatMembersCount(tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID})
		if err != nil {
			log.Fatal(err)
		}

		command := update.Message.Command()
		switch command {
		case "judgeList":
			phrases := commands.JudgeList()
			answer := ""
			for _, judgePhrase := range phrases {
				// N - chatMembersCount, x - in favor, y - against
				// x - y >= N / 3 && x >= N / 2
				prefix := "- "
				if inFavor(judgePhrase, chatMembersCount) {
					prefix = "+ "
				}
				answer +=  prefix + judgePhrase.Phrase + " " + strconv.Itoa(judgePhrase.Voteup) + " " + strconv.Itoa(judgePhrase.Votedown) + "\n"
			}
			message := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			bot.Send(message)

		case "judgeAdd":
			commands.JudgeVote(update.Message.From.ID, update.Message.CommandArguments(), true)

		case "judgeRemove":
			commands.JudgeVote(update.Message.From.ID, update.Message.CommandArguments(), false)
		}
	}
}
func inFavor(judgePhrase database.JudgePhraseInfo, chatMembersCount int) bool {
	return float64(judgePhrase.Voteup-judgePhrase.Votedown) >= float64(chatMembersCount)/3 && float64(judgePhrase.Voteup+judgePhrase.Votedown) >= float64(chatMembersCount)/2
}
