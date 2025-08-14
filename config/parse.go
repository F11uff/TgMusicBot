package config

import (
	"github.com/spf13/viper"
	"log"
)

func ParseConfigDatabase(cnf *Config) {
	viper.SetConfigName("db")
	viper.AddConfigPath("config/.")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Error reading config file, ", err)
	}

	cnf.Database.DatabaseName = viper.GetString("Database.database_name")
	cnf.Database.Username = viper.GetString("Database.username")
	cnf.Database.Port = viper.GetString("Database.port")
	cnf.Database.SslMode = viper.GetString("Database.sslmode")
	cnf.Database.Host = viper.GetString("Database.host")

	cnf.Database.Password = viper.GetInt("Database.password")
}
