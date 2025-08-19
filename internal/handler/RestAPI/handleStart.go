package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/model"
	"musicBot/internal/storage"
)

func HandleStart(md *model.Model, db *storage.Database, msg *tgbotapi.Message) error {
	reply := HandleStartCommand(msg.Chat.ID)
	reply.ReplyMarkup = createMainKeyboard()
	_, err := md.Bot.Send(reply)

	TGUsername := msg.From.UserName

	if err := db.AddUserRequest(TGUsername); err != nil {
		return err
	}

	return err
}
