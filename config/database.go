package config

type Database struct {
	Host         string `yml:"host" env-default:""`
	Port         string `yml:"port" env-default:""`
	Username     string `yml:"username" env-default:""`
	Password     int    `yml:"password" env-default:""`
	DatabaseName string `yml:"database_name" env-default:""`
	SslMode      string `yml:"sslMode" env-default:""`
}

func NewDatabase() *Database {
	return &Database{
		Host:         "",
		Port:         "",
		Username:     "",
		Password:     0,
		DatabaseName: "",
		SslMode:      "",
	}
}
