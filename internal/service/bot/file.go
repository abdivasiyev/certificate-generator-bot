package bot

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *service) sendPhoto(ctx context.Context, msg SendPhoto) (Message, error) {
	var (
		params = map[string]any{}
	)

	bytes, err := json.Marshal(msg)
	if err != nil {
		return Message{}, err
	}

	if err = json.Unmarshal(bytes, &params); err != nil {
		return Message{}, err
	}

	delete(params, "photo")

	response, err := sendFile[Message](
		ctx,
		s.httpClient,
		"/sendPhoto",
		params,
		"photo",
		msg.Photo,
	)
	if err != nil {
		return Message{}, fmt.Errorf("sendPhoto: %w", err)
	}

	if !response.Ok {
		return Message{}, fmt.Errorf("sendPhoto: %s", response.Description)
	}

	return response.Result, nil
}
