package service

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"musicBot/internal/model"
	"regexp"
	"strings"
)

func ParseArtistTitle(music *model.Music, msg *tgbotapi.Message) error {
	req := strings.TrimSpace(msg.Text)

	if !strings.Contains(req, ":") {
		musicClient := strings.TrimSpace(msg.Text)
		music = music.SetMusic(musicClient)
	} else {
		re := regexp.MustCompile(`\s*:\s*`)
		parts := re.Split(req, 2)

		if len(parts) != 2 {
			return errors.New("invalid artist title")
		}

		artist := strings.TrimSpace(parts[0])
		title := strings.TrimSpace(parts[1])

		if artist == "" || title == "" {
			return errors.New("artist title or artist or music title empty")
		}

		music = music.SetMusic(title)
		music = music.SetArtist(artist)
	}

	return nil
}
