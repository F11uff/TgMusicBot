package service

import "github.com/joho/godotenv"

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		// logger
	}
}
