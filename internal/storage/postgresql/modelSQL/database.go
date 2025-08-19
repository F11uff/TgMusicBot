package modelSQL

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"musicBot/config"
	"time"
)

type PosgreSQLDatabase struct {
	psql  *sql.DB
	Music *Music
	//User  *User
}

func NewPosgreSQLDatabase() *PosgreSQLDatabase {
	var psql2 *sql.DB

	db := &PosgreSQLDatabase{
		psql:  psql2,
		Music: NewMusic(),
		//User:  NewUser(),
	}

	return db
}

func (db *PosgreSQLDatabase) ConnectionURL(config *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%d dbname=%s sslmode=%s ",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.DatabaseName,
		config.Database.SslMode,
	)
}

func (db *PosgreSQLDatabase) Connect(urlConnect string) error {
	var err error

	db.psql, err = sql.Open("postgres", urlConnect)

	if err != nil {
		return errors.New("fail to connection to database")
	}

	if err = db.psql.Ping(); err != nil {
		return errors.New("fail to ping database")
	}

	return nil
}

func (db *PosgreSQLDatabase) Close() error {
	if db.psql != nil {
		return db.psql.Close()
	}

	return nil
}

func (db *PosgreSQLDatabase) GetLikedSongRequest() ([]Music, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sqlRequest := `SELECT music, artist FROM LikeMusic`

	rows, err := db.psql.QueryContext(ctx, sqlRequest)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musicList []Music

	for rows.Next() {
		var music, artist string

		if err := rows.Scan(&music, &artist); err != nil {
			return nil, err
		}

		musicList = append(musicList, *NewMusic().SetArtist(artist).SetTitle(music))
	}

	return musicList, nil
}

func (db *PosgreSQLDatabase) AddLikedSongRequest(artist, title string) error {
	sqlRequest := `INSERT INTO LikeMusic(artist, music) VALUES ($1, $2)`

	_, err := db.psql.Exec(sqlRequest, artist, title)

	if err != nil {
		return err
	}

	return nil
}

func (db *PosgreSQLDatabase) RemoveLikedSongRequest() error {
	return nil
}

func (db *PosgreSQLDatabase) AddUserRequest(username string) error {
	sqlRequest := `INSERT INTO users(username) VALUES ($1)`

	_, err := db.psql.Exec(sqlRequest, username)

	if err != nil {
		return errors.New("fail to add user to database")
	}

	return nil
}
