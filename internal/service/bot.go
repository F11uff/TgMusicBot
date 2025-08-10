package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	"musicBot/internal/handler/RestAPI"
	"musicBot/internal/model"
)

func Endpoints(channel tgbotapi.UpdatesChannel, config *config.Config, user *model.User) {
	for update := range channel {
		if update.Message != nil {
			RestAPI.HandleMessage(config, user, update.Message)
		}

		if update.CallbackQuery != nil {
			RestAPI.HandleCallback(config, user, update.CallbackQuery)
		}
	}
}
