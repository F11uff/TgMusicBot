package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	Bot           *tgbotapi.BotAPI
	youtubeAPIKey string
}

func (conf *Config) InitConfig() *Config {
	return &Config{
		Bot:           nil,
		youtubeAPIKey: "",
	}
}

func (conf *Config) InitBot(bot *tgbotapi.BotAPI) *Config {
	conf.Bot = bot
	return conf
}

func (cnf *Config) SetYoutubeAPIKey(APIKey string) *Config {
	cnf.youtubeAPIKey = APIKey

	return cnf
}

func (cnf *Config) GetYoutubeAPIKey() string {
	return cnf.youtubeAPIKey
}
