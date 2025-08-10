package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/model"
	"musicBot/internal/service"
	"musicBot/pkg"
	"os"
)

func main() {
	pkg.InitEnv()

	var bot *tgbotapi.BotAPI
	var conf *config.Config

	conf = conf.InitConfig()

	token := os.Getenv("NEW_BOT_TOKEN")
	if token == "" {
		log.Fatal("NEW_BOT_TOKEN невозможно найти, возможно, вы запускаете не с корневой папки")
	}

	user := model.NewUser()

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		//сюда логгер написать надо(собственный) или взять мой логгер с моего гита
	}

	updateBot := tgbotapi.NewUpdate(0)
	updateBot.Timeout = 45

	conf = conf.InitBot(bot)

	conf.Bot.Debug = true

	updates := conf.Bot.GetUpdatesChan(updateBot)

	service.Endpoints(updates, conf, user)
}
