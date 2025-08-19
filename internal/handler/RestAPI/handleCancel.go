package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/model"
)

func HandleCancel(md *model.Model, msg *tgbotapi.Message) error {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "Хорошо")
	reply.ReplyMarkup = createMainKeyboard()
	md.Bot.Send(reply)

	return nil
}
