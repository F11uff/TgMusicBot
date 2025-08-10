package RestAPI

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/model"
)

const STATE = "Wait"

func HandleCallback(conf *config.Config, user *model.User, query *tgbotapi.CallbackQuery) {
	if _, err := conf.Bot.Request(tgbotapi.NewCallback(query.ID, "☻")); err != nil {
		//log.Printf("Callback error: %v", err)
		return
	}

	var response string

	fmt.Println("________________________________", query.From.ID)
	fmt.Println("________________________________", query.From.UserName)

	switch query.Data {
	case "search":
		user.SetUserState(query.From.ID, STATE)

		response = "Введите имя исполнителя и название песни для воспроизведения(Пример - 'SLAVA SKRIPKA - Бобр'): "

	}

	msg := tgbotapi.NewMessage(query.Message.Chat.ID, response)
	msg.ParseMode = "Markdown"
	if _, err := conf.Bot.Send(msg); err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
	}
}
