package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Model struct {
	Bot   *tgbotapi.BotAPI
	Music *Music
	User  *User
}

func NewModel(bot *tgbotapi.BotAPI) *Model {
	return &Model{
		bot,
		NewMusic(),
		NewUser(),
	}
}
