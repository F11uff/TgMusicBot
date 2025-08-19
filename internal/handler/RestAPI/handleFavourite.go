package RestAPI

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/model"
	"musicBot/internal/storage"
)

func HandleFavourite(md *model.Model, db *storage.Database, msg *tgbotapi.Message) error {
	_, _ = md.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Ваши избранные треки:"))

	musicList, err := db.GetLikedSongRequest()

	if err != nil {
		return err
	}

	for index, music := range musicList {
		message := fmt.Sprintf("%d) %s - %s", index, music.GetArtist(), music.GetTitle())

		_, _ = md.Bot.Send(tgbotapi.NewMessage(msg.Chat.ID, message))
	}

	return nil
}
