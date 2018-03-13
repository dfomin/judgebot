package network

import (
	"judgebot/commands"
	"judgebot/private"
	"log"
	"net/http"
	"strings"

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
		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		chatMembersCount, err := bot.GetChatMembersCount(tgbotapi.ChatConfig{ChatID: chatID})
		if err != nil {
			log.Fatal(err)
		}

		chatMembersCount -= 1

		command := strings.ToLower(update.Message.Command())
		args := strings.Trim(update.Message.CommandArguments(), " ")
		switch command {
		case "judgelist":
			answer := commands.JudgeList(chatID, chatMembersCount)
			message := tgbotapi.NewMessage(chatID, answer)
			bot.Send(message)

		case "judgeadd":
			if len(args) != 0 {
				commands.JudgeVote(update.Message.From.ID, chatID, args, true)
			}

		case "judgeremove":
			if len(args) != 0 {
				commands.JudgeVote(update.Message.From.ID, chatID, args, false)
			}

		case "judge":
			names := strings.Split(args, " ")
			if len(names) > 0 {
				answer := commands.Judge(names, chatID, chatMembersCount)
				message := tgbotapi.NewMessage(chatID, answer)
				bot.Send(message)
			}
		}
	}
}
