package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	"musicBot/internal/handler/RestAPI"
	"musicBot/internal/model"
)

func Endpoints(channel tgbotapi.UpdatesChannel, config *config.Config, user *model.User) error {
	var err error

	for update := range channel {
		if update.Message != nil {
			err = RestAPI.HandleMessage(config, user, update.Message)
			if err != nil {

				return err
			}
		}

		//if update.CallbackQuery != nil {
		//	RestAPI.HandleCallback(config, user, update.CallbackQuery)
		//}
	}

	return nil
}
