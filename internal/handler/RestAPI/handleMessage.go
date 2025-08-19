package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	_const "musicBot/internal/const"
	"musicBot/internal/model"
	"musicBot/internal/storage"
)

func HandleMessage(conf *config.Config, md *model.Model, db *storage.Database, msg *tgbotapi.Message) error {
	switch {
	case msg.IsCommand():
		switch msg.Command() {
		case "start":
			return HandleStart(md, db, msg)
		}
	case msg.Text == "🔍 Поиск":
		return HandleSearch(md, msg)
	case msg.Text == "📁 Избранное":
		return HandleFavourite(md, db, msg)
	case msg.Text == "❤️ Добавить в избранное":
		return HandleAddFavourite(md, db, msg)
	case msg.Text == "❌ Отмена":
		return HandleCancel(md, msg)

	default:
		if state, ok := md.User.GetUserState(msg.From.ID); ok && state == _const.STATE {
			return HandleParseMusic(md, conf, msg)
		}
	}

	return nil
}

func createMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔍 Поиск"),
			tgbotapi.NewKeyboardButton("📁 Избранное"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🎧 Плейлисты", "playlists"),
		//	tgbotapi.NewInlineKeyboardButtonData("⚙️ Настройки", "settings"),
		//),
	)
}

func createAddKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("❤️ Добавить в избранное"),
			tgbotapi.NewKeyboardButton("❌ Отмена"),
		),
	)
}
