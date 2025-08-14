package RestAPI

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/config"
	"musicBot/internal/model"
	"musicBot/internal/service"
)

func HandleMusicRequest(cnf *config.Config, music *model.Music, msg *tgbotapi.Message) error {
	query := music.GetMusic()

	videoURL, err := service.SearchMusic(cnf.GetYoutubeAPIKey(), query)

	if err != nil {
		return err
	}

	reply := tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(
		"По вашему запросу найдено: *%s* — %s\n▶️ [Слушать на YouTube](%s)",
		music.GetArtist(), music.GetMusic(), videoURL))

	reply.ParseMode = "Markdown"
	cnf.Bot.Send(reply)

	return err
}
