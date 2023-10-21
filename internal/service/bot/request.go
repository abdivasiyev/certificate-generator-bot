package bot

import (
	"context"
	"fmt"

	"github.com/imroc/req/v3"
)

func sendRequest[T any](ctx context.Context,
	client *req.Client,
	action string,
	params any,
) (Response[T], error) {
	var response Response[T]

	r := client.Post("/bot{token}" + action).
		SetBody(params).
		Do(ctx)

	if r == nil {
		return response, fmt.Errorf("sendRequest: %w", ErrTelegramRequest)
	}

	if err := r.Unmarshal(&response); err != nil {
		return response, fmt.Errorf("sendRequest: %w", err)
	}

	return response, nil
}
