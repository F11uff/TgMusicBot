package config

type Config struct {
	Database      *Database `yml:"Database"`
	youtubeAPIKey string
}

func (conf *Config) InitConfig() *Config {
	return &Config{
		Database:      nil,
		youtubeAPIKey: "",
	}
}

func (cnf *Config) SetYoutubeAPIKey(APIKey string) *Config {
	cnf.youtubeAPIKey = APIKey

	return cnf
}

func (cnf *Config) GetYoutubeAPIKey() string {
	return cnf.youtubeAPIKey
}
