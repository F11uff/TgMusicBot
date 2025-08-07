package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/service"
	"os"
)

func main() {
	service.InitEnv()

	var bot *tgbotapi.BotAPI
	var config config.Config

	config.InitConfig()

	token := os.Getenv("NEW_BOT_TOKEN")
	if token == "" {
		log.Fatal("NEW_BOT_TOKEN невозможно найти, возможно, вы запускаете не с корневой папки")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		//сюда логгер написать надо(собственный) или взять мой логгер с моего гита
	}

	config.Bot = bot

	config.Bot.Debug = true

	updateBot := tgbotapi.NewUpdate(0)
	updateBot.Timeout = 45

	updates := config.Bot.GetUpdatesChan(updateBot)

	service.Endpoints(updates, config)
}
