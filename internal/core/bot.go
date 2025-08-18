package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	"musicBot/internal/handler/RestAPI"
	"musicBot/internal/model"
	"musicBot/internal/storage"
	"musicBot/internal/storage/postgresql"
)

func Endpoints(channel tgbotapi.UpdatesChannel, conf *config.Config, md *model.Model, db *storage.Database) error {
	var err error

	go func() {
		postgresql.ConnDB(conf, db)
	}()

	for update := range channel {
		if update.Message != nil {
			err = RestAPI.HandleMessage(conf, md, db, update.Message)
			if err != nil {

				return err
			}
		}
	}

	defer db.Close()

	return nil
}
