package bot

import (
	"context"
	"fmt"
)

func (s *service) setWebhook(ctx context.Context) error {
	if err := s.deleteWebhook(ctx); err != nil {
		return err
	}

	response, err := sendRequest[bool](ctx, s.httpClient, "/setWebhook", SetWebhook{
		URL:         s.webhookURL + "/webhook/" + s.botToken,
		SecretToken: s.webhookSecret,
	})
	if err != nil {
		return fmt.Errorf("setWebhook: %w", err)
	}

	if !response.Ok || !response.Result {
		s.logger.Error().Str("setWebhook error description", response.Description).Err(err).Send()
		return fmt.Errorf("setWebhook: %w", ErrSetWebhook)
	}

	return nil
}

func (s *service) deleteWebhook(ctx context.Context) error {
	response, err := sendRequest[bool](ctx, s.httpClient, "/deleteWebhook", DeleteWebhook{
		DropPendingUpdates: false,
	})
	if err != nil {
		return fmt.Errorf("deleteWebhook: %w", err)
	}

	if !response.Ok || !response.Result {
		s.logger.Error().Str("deleteWebhook error description", response.Description).Err(err).Send()
		return fmt.Errorf("deleteWebhook: %w", ErrDeleteWebhook)
	}

	return nil
}
