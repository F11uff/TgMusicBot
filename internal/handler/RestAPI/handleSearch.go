package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_const "musicBot/internal/const"
	"musicBot/internal/model"
)

func HandleSearch(md *model.Model, msg *tgbotapi.Message) error {
	md.User.SetUserState(msg.From.ID, _const.STATE)

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Введите имя исполнителя и название песни для воспроизведения(Пример - \\\"SLAVA SKRIPKA : Бобр\\\" или просто название песни): \"")
	reply.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

	_, err := md.Bot.Send(reply)

	return err
}
