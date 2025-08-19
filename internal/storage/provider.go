package storage

import (
	"musicBot/config"
	"musicBot/internal/model"
)

type MusicProvider interface {
	Connect(urlConnect string) error
	ConnectionURL(config *config.Config) string

	GetLikedSongRequest() ([]model.Music, error)

	Close() error
}

type Database struct {
	music MusicProvider
}

func NewDatabase(music MusicProvider) *Database {
	return &Database{music: music}
}

func (db *Database) GetLikedSong() ([]model.Music, error) {
	return db.music.GetLikedSongRequest()
}

func (db *Database) Connect(urlConnect string) error {
	return db.music.Connect(urlConnect)
}

func (db *Database) Close() error {
	return db.music.Close()
}

func (db *Database) ConnectionURL(config *config.Config) string {
	return db.music.ConnectionURL(config)
}
