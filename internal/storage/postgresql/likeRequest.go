package postgresql

import (
	"context"
	"musicBot/internal/model"
	"musicBot/internal/storage"
	"time"
)

func LikeRequest(db *storage.Database) ([]model.Music, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sqlRequest := `SELECT music, artist FROM LikeMusic`

	rows, err := db.GetDB().QueryContext(ctx, sqlRequest)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var musicList []model.Music

	for rows.Next() {
		var music, artist string

		if err := rows.Scan(&music, &artist); err != nil {
			return nil, err
		}

		musicList = append(musicList, *model.NewMusic().SetArtist(artist).SetMusic(music))
	}

	return musicList, nil
}
