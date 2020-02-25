package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	webHook = "" 
)

func main()  {
	token := os.Getenv("TOKEN")
	port := os.Getenv("PORT")

	go func ()  {
		_ = http.ListenAndServe(":" + port, nil)
	}()

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Create bot error: ", err)
	}
	log.Println("Bot created")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("Setting webHook %v: error: %v", webHook, err)
	}
	log.Println("webHook setting set")

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
			log.Println(err)
		}
	}
}