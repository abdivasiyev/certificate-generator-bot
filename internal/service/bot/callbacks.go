package bot

import "context"

func (s *service) handleStartTest(ctx context.Context, callback *CallbackQuery) error {
	_, err := s.editMessage(ctx, EditMessage{
		ChatID:    callback.Message.Chat.ID,
		MessageID: callback.Message.MessageID,
		Text:      "Test boshlandi!",
		ParseMode: "html",
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	return nil
}
