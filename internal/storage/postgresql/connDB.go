package postgresql

import (
	"fmt"
	"musicBot/config"
	"musicBot/internal/storage"
)

func ConnDB(conf *config.Config, db *storage.Database) (*storage.Database, error) {
	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%d dbname=%s sslmode=%s ",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.DatabaseName,
		conf.Database.SslMode,
	)

	var err error

	db, err = db.Connect(urlConnection)

	if err != nil {

		return nil, err
	}

	return db, nil
}
