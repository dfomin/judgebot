package network

import (
	"judgebot/commands"
	"judgebot/private"
	"log"
	"net/http"

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
		chatMembersCount, err := bot.GetChatMembersCount(tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID})
		if err != nil {
			log.Fatal(err)
		}

		command := update.Message.Command()
		switch command {
		case "judgeList":
			answer := commands.JudgeList(chatMembersCount)
			message := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			bot.Send(message)

		case "judgeAdd":
			args := update.Message.CommandArguments()
			if len(args) != 0 {
				commands.JudgeVote(update.Message.From.ID, args, true)
			}

		case "judgeRemove":
			args := update.Message.CommandArguments()
			if len(args) != 0 {
				commands.JudgeVote(update.Message.From.ID, args, false)
			}
		}
	}
}
