package bot

import "context"

func (s *service) Handle(ctx context.Context, update Update) error {
	s.logger.Debug().Any("update from telegram", update).Send()

	msg, err := s.sendMessage(ctx, SendMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      update.Message.Text,
		ParseMode: "html",
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	s.logger.Debug().Any("message", msg).Send()

	return nil
}
