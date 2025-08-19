package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/model"
	"musicBot/internal/storage"
)

func HandleAddFavourite(md *model.Model, db *storage.Database, msg *tgbotapi.Message) error {
	err := db.AddLikedSongRequest(md.Music.GetArtist(), md.Music.GetMusic())

	if err != nil {
		return err
	}
	md.Music.ClearArtistAndMusic()

	reply := tgbotapi.NewMessage(msg.Chat.ID, "Хорошо")
	reply.ReplyMarkup = createMainKeyboard()
	md.Bot.Send(reply)

	return nil
}
