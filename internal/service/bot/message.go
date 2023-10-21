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

	if !response.Ok {
		return Message{}, fmt.Errorf("sendMessage: %s", response.Description)
	}

	return response.Result, nil
}

func (s *service) editMessage(ctx context.Context, msg EditMessage) (Message, error) {
	response, err := sendRequest[Message](ctx, s.httpClient, "/editMessageText", msg)
	if err != nil {
		return Message{}, fmt.Errorf("editMessage: %w", err)
	}

	if !response.Ok {
		return Message{}, fmt.Errorf("editMessage: %s", response.Description)
	}

	return response.Result, nil
}

func (s *service) editReplyMarkup(ctx context.Context, msg EditReplyMarkup) (Message, error) {
	response, err := sendRequest[Message](ctx, s.httpClient, "/editMessageReplyMarkup", msg)
	if err != nil {
		return Message{}, fmt.Errorf("editMessage: %w", err)
	}

	if !response.Ok {
		return Message{}, fmt.Errorf("editMessage: %s", response.Description)
	}

	return response.Result, nil
}
