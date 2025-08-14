package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	_const "musicBot/internal/const"
	"musicBot/internal/model"
	"musicBot/internal/service"
)

func HandleMessage(conf *config.Config, user *model.User, msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID

	switch {
	case msg.IsCommand():
		switch msg.Command() {
		case "start":
			reply := HandleStartCommand(chatID)
			reply.ReplyMarkup = createMainKeyboard()
			_, err := conf.Bot.Send(reply)
			return err
		}
	case msg.Text == "🔍 Поиск":
		user.SetUserState(msg.From.ID, _const.STATE)

		reply := tgbotapi.NewMessage(chatID, "Введите имя исполнителя и название песни для воспроизведения(Пример - \\\"SLAVA SKRIPKA - Бобр\\\" или просто название песни): \"")
		reply.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

		_, err := conf.Bot.Send(reply)

		return err
	case msg.Text == "📁 История музыки":
		_, err := conf.Bot.Send(tgbotapi.NewMessage(chatID, "История прослушивания..."))

		return err
	default:
		if state, ok := user.GetUserState(msg.From.ID); ok && state == _const.STATE {
			music := model.NewMusic("", "")
			err := service.ParseArtistTitle(music, msg)

			if err != nil {
				return err
			}

			err = HandleMusicRequest(conf, music, msg)
			if err != nil {
				return err
			}

			music.ClearArtistAndMusic()
			user.ClearUserState(msg.From.ID)

			reply := tgbotapi.NewMessage(chatID, "Готово! Выберите следующее действие")
			reply.ReplyMarkup = createMainKeyboard()
			_, err = conf.Bot.Send(reply)

			return err
		}
	}

	return nil
}

func createMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🔍 Поиск"),
			tgbotapi.NewKeyboardButton("📁 История музыки"),
		),
		//tgbotapi.NewInlineKeyboardRow(
		//	tgbotapi.NewInlineKeyboardButtonData("🎧 Плейлисты", "playlists"),
		//	tgbotapi.NewInlineKeyboardButtonData("⚙️ Настройки", "settings"),
		//),
	)
}
