package RestAPI

//
//import (
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	"log"
//	"musicBot/config"
//	_const "musicBot/internal/const"
//	"musicBot/internal/modelSQL"
//)
//
//func HandleCallback(conf *config.Config, user *modelSQL.User, query *tgbotapi.CallbackQuery) {
//	if _, err := conf.Bot.Request(tgbotapi.NewCallback(query.ID, "☻")); err != nil {
//		log.Printf("Callback error: %v", err)
//	}
//
//	var response string
//
//	switch query.Data {
//	case "search":
//		user.SetUserState(query.From.ID, _const.STATE)
//
//		response = "Введите имя исполнителя и название песни для воспроизведения(Пример - \"SLAVA SKRIPKA - Бобр\" или просто название песни): "
//
//	}
//
//	msg := tgbotapi.NewMessage(query.Message.Chat.ID, response)
//	msg.ParseMode = "Markdown"
//	if _, err := conf.Bot.Send(msg); err != nil {
//		log.Printf("Ошибка отправки сообщения: %v", err)
//	}
//}
