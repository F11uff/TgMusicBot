package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	_const "musicBot/internal/const"
)

func Endpoints(channel tgbotapi.UpdatesChannel, config config.Config) {
	for update := range channel {
		if update.Message != nil {
			var msg tgbotapi.MessageConfig

			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, с помощью меня вы сможете слушать музыку в телеграмме")
				//msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы нажали на кнопку!")
				//msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
				//
				//keyboard := tgbotapi.NewReplyKeyboard(
				//	tgbotapi.NewKeyboardButtonRow(
				//		tgbotapi.NewKeyboardButton("Нажми меня"),
				//	),
				//)
				//msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Нажмите на кнопку:")
				//msg.ReplyMarkup = keyboard
				//config.Bot.Send(msg)
			}

			config.Bot.Send(msg)
		}
	}
}

func createLabel(messageConfig tgbotapi.MessageConfig) {
	messageConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	for str, _ := range _const.GetBillet() {

	}
}
