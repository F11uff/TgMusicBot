package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	"musicBot/internal/model"
)

func HandleMessage(conf *config.Config, user *model.User, msg *tgbotapi.Message) {
	chatID := msg.Chat.ID

	switch {
	case msg.IsCommand():
		switch msg.Command() {
		case "start":
			reply := HandleStartCommand(chatID)
			reply.ReplyMarkup = createMainKeyboard()
			SendMessage(conf.Bot, reply)
		}
	default:
		if _, ok := user.GetUserState(msg.From.ID); ok {

		}
	}
}

func SendMessage(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	if _, err := bot.Send(msg); err != nil {
		//log.Printf("Ошибка отправки сообщения: %v", err)
	}
}

func createMainKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔍 Поиск", "search"),
			//tgbotapi.NewInlineKeyboardButtonData("📁 Моя музыка", "my_music"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🎧 Плейлисты", "playlists"),
		//	tgbotapi.NewInlineKeyboardButtonData("⚙️ Настройки", "settings"),
		//),
	)
}
