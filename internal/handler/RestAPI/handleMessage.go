package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	_const "musicBot/internal/const"
	"musicBot/internal/model"
	"musicBot/pkg"
)

func HandleMessage(conf *config.Config, user *model.User, msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID

	switch {
	case msg.IsCommand():
		switch msg.Command() {
		case "start":
			reply := HandleStartCommand(chatID)
			reply.ReplyMarkup = createMainKeyboard()
			conf.Bot.Send(reply)
		}
	default:
		if state, ok := user.GetUserState(msg.From.ID); ok {
			switch state {
			case _const.STATE:

				music := model.NewMusic("", "")
				err := pkg.ParseArtistTitle(music, msg)

				if err != nil {
					return err
				}

				err = HandleMusicRequest(conf, music, msg)

				if err != nil {
					return err
				}

				music.ClearArtistAndMusic()
				user.ClearUserState(msg.From.ID)
			default:

			}

		}
	}

	return nil
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
