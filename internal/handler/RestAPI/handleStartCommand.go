package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartCommand(chatID int64) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, "Привет, с помощью меня вы сможете слушать музыку в телеграмме")
}
