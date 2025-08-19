package storage

import (
	"musicBot/config"
	"musicBot/internal/storage/postgresql/modelSQL"
)

type MusicProvider interface {
	//Connect to DB
	Connect(urlConnect string) error
	ConnectionURL(config *config.Config) string

	//Methods for songs
	GetLikedSongRequest() ([]modelSQL.Music, error)
	AddLikedSongRequest(artist, title string) error
	RemoveLikedSongRequest() error

	//Methods for user
	AddUserRequest(username string) error

	//Close DB
	Close() error
}

type Database struct {
	music MusicProvider
}

func NewDatabase(music MusicProvider) *Database {
	return &Database{
		music: music,
	}
}

func (db *Database) GetLikedSongRequest() ([]modelSQL.Music, error) {
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

func (db *Database) AddUserRequest(username string) error {
	return db.music.AddUserRequest(username)
}

func (db *Database) AddLikedSongRequest(artist, title string) error {
	return db.music.AddLikedSongRequest(artist, title)
}
