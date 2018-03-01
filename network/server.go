package network

import (
	"judgebot/commands"
	"judgebot/private"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
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
		command := update.Message.Command()
		switch command {
		case "judgeList":
			phrases := commands.JudgeList()
			answer := ""
			for _, judgePhrase := range phrases {
				answer += judgePhrase.Phrase + " " + strconv.Itoa(judgePhrase.Voteup) + " " + strconv.Itoa(judgePhrase.Votedown) + "\n"
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
