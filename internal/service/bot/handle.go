package bot

import (
	"context"
	"strings"
)

func (s *service) Handle(ctx context.Context, update Update) error {
	s.logger.Debug().Any("update from telegram", update).Send()

	if err := s.handleTextCommands(ctx, update); err != nil {
		s.logger.Error().Err(err).Msg("failed to handle text command")
		return err
	}

	if err := s.handleCallbackCommands(ctx, update); err != nil {
		s.logger.Error().Err(err).Msg("failed to handle text command")
		return err
	}

	return nil
}

func (s *service) handleCallbackCommands(ctx context.Context, update Update) error {
	if update.CallbackQuery == nil {
		return nil
	}

	switch strings.TrimSpace(update.CallbackQuery.Data) {
	case "/start_test":
		return s.handleStartTest(ctx, update.CallbackQuery)
	case "/option_a", "/option_b", "/option_c", "/option_d":
		return s.handleTestOption(ctx, update.CallbackQuery)
	case "/generate_certificate":
		return s.handleGenerateCertificate(ctx, update.CallbackQuery)
	default:
		return s.handleUnknown(ctx, update)
	}
}

func (s *service) handleTextCommands(ctx context.Context, update Update) error {
	if update.Message == nil {
		return nil
	}

	switch strings.TrimSpace(update.Message.Text) {
	case "/start":
		return s.handleStart(ctx, update.Message)
	default:
		return s.handleUnknown(ctx, update)
	}
}
