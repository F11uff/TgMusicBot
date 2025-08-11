package service

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func SearchMusic(apiKey string, query string) (string, error) {
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		return "", errors.New("YouTube API Error")
	}

	call := service.Search.List([]string{"id,snippet"}).Q(query).MaxResults(1).Type("Video")

	response, err := call.Do()

	if err != nil {
		return "", errors.New("Fail Search Music")
	}

	if len(response.Items) == 0 {
		return "", errors.New("Fail Search Music")
	}

	video := response.Items[0]

	return fmt.Sprintf("https://youtu.be/%s", video.Id.VideoId), nil
}
