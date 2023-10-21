package bot

import (
	"context"
	"fmt"
)

func (s *service) sendMessage(ctx context.Context, msg SendMessage) (Message, error) {
	response, err := sendRequest[Message](ctx, s.httpClient, "/sendMessage", msg)
	if err != nil {
		return Message{}, fmt.Errorf("sendMessage: %w", err)
	}

	return response.Result, nil
}
