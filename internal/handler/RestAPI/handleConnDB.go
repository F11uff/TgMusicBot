package RestAPI

import (
	"database/sql"
	"fmt"
	"musicBot/config"
	"musicBot/internal/storage/postgresql"
)

var psql *sql.DB

func handleConnDB(conf *config.Config) error {
	db := postgresql.NewDatabase(psql)

	urlConnection := fmt.Sprintf("host=%s port=%s username=%s password=%d dbname=%s sslmode=%s ",
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
		return err
	}

	return nil
}
