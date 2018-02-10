package network

import (
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

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://pigowl.com:88/"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/")
	go http.ListenAndServeTLS(":88", "fullchain.pem", "privkey.pem", nil)

	for update := range updates {
		command := update.Message.Command()
		switch command {
		case "start":
		}
	}
}
