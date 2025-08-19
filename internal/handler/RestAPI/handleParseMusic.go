package RestAPI

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"musicBot/config"
	"musicBot/internal/model"
	"musicBot/internal/service"
	"time"
)

func HandleParseMusic(md *model.Model, conf *config.Config, msg *tgbotapi.Message) error {
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

	md.User.ClearUserState(msg.From.ID)

	reply := tgbotapi.NewMessage(msg.From.ID, "Готово! Выберите следующее действие")

	reply.ReplyMarkup = createAddKeyboard()

	_, _ = md.Bot.Send(reply)

	return nil
}
