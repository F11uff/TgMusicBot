package config

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Config struct {
	Bot *tgbotapi.BotAPI
}

func (conf *Config) InitConfig() *Config {
	return &Config{
		Bot: nil,
	}
}
