package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/core"
	"musicBot/internal/model"
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

	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")

	if youtubeAPIKey == "" {
		log.Fatal("Не найден YOUTUBE_API_KEY")
	}

	user := model.NewUser()

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		//сюда логгер написать надо(собственный) или взять мой логгер с моего гита
	}

	updateBot := tgbotapi.NewUpdate(0)
	updateBot.Timeout = 45

	conf = conf.InitBot(bot)
	conf = conf.SetYoutubeAPIKey(youtubeAPIKey)

	conf.Bot.Debug = true

	updates := conf.Bot.GetUpdatesChan(updateBot)

	err = core.Endpoints(updates, conf, user)

	if err != nil {
		log.Fatal(err)
	}
}
