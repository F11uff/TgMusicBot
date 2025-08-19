package RestAPI

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	_const "musicBot/internal/const"
	"musicBot/internal/model"
	"musicBot/internal/service"
	"musicBot/internal/storage"
	"time"
)

func HandleMessage(conf *config.Config, md *model.Model, db *storage.Database, msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID

	switch {
	case msg.IsCommand():
		switch msg.Command() {
		case "start":
			reply := HandleStartCommand(chatID)
			reply.ReplyMarkup = createMainKeyboard()
			_, err := md.Bot.Send(reply)
			return err
		}
	case msg.Text == "🔍 Поиск":
		md.User.SetUserState(msg.From.ID, _const.STATE)

		reply := tgbotapi.NewMessage(chatID, "Введите имя исполнителя и название песни для воспроизведения(Пример - \\\"SLAVA SKRIPKA : Бобр\\\" или просто название песни): \"")
		reply.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)

		_, err := md.Bot.Send(reply)

		return err
	case msg.Text == "📁 Избранное":
		_, err := md.Bot.Send(tgbotapi.NewMessage(chatID, "Ваши избранные треки:"))

		musicList, err := db.GetLikedSong()

		for index, music := range musicList {
			message := fmt.Sprintf("%d) %s - %s", index, music.GetArtist(), music.GetMusic())

			_, _ = md.Bot.Send(tgbotapi.NewMessage(chatID, message))
		}

		return err
	default:
		if state, ok := md.User.GetUserState(msg.From.ID); ok && state == _const.STATE {
			err := service.ParseArtistTitle(md.Music, msg)

			modelCopy := *md
			msgCopy := *msg

			if err != nil {
				return err
			}

			go func() {
				err = HandleMusicRequest(conf, &modelCopy, &msgCopy)

				log.Print(err)
			}()

			time.Sleep(1 * time.Second)

			md.Music.ClearArtistAndMusic()
			md.User.ClearUserState(msg.From.ID)

			reply := tgbotapi.NewMessage(chatID, "Готово! Выберите следующее действие")
			reply.ReplyMarkup = createMainKeyboard()
			_, err = md.Bot.Send(reply)

			return err
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
