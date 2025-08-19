package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/core"
	"musicBot/internal/model"
	"musicBot/internal/storage"
	model2 "musicBot/internal/storage/postgresql/model"
	"musicBot/pkg"
	"os"
)

func main() {
	pkg.InitEnv()

	var bot *tgbotapi.BotAPI
	var conf *config.Config

	conf = conf.InitConfig()
	conf.Database = config.NewDatabase()
	config.ParseConfigDatabase(conf)

	//fmt.Println(conf.Database.DatabaseName, " ", conf.Database.Password, " ", conf.Database.SslMode, " ", conf.Database.Port, " ", conf.Database.Host, " ", conf.Database.Username)

	token := os.Getenv("NEW_BOT_TOKEN")
	if token == "" {
		log.Fatal("NEW_BOT_TOKEN невозможно найти, возможно, вы запускаете не с корневой папки")
	}

	youtubeAPIKey := os.Getenv("YOUTUBE_API_KEY")

	if youtubeAPIKey == "" {
		log.Fatal("Не найден YOUTUBE_API_KEY")
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Println(err)
	}

	updateBot := tgbotapi.NewUpdate(0)
	updateBot.Timeout = 45

	md := model.NewModel(bot)

	db := model2.NewPosgreSQLDatabase()
	musicRepo := storage.NewDatabase(db)

	conf = conf.SetYoutubeAPIKey(youtubeAPIKey)

	updates := md.Bot.GetUpdatesChan(updateBot)

	md.Bot.Debug = true

	err = core.Endpoints(updates, conf, md, musicRepo)

	if err != nil {
		log.Println(err)
	}
}
