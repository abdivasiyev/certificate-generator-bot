package bot

import "context"

func (s *service) handleUnknown(ctx context.Context, update Update) error {
	_, err := s.sendMessage(ctx, SendMessage{
		ChatID:    update.Message.Chat.ID,
		Text:      "Unknown command",
		ParseMode: "html",
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	return nil
}

func (s *service) handleStart(ctx context.Context, message *Message) error {
	s.logger.Debug().Any("message", message).Msg("handling start from user")
	msg, err := s.sendMessage(ctx, SendMessage{
		ChatID: message.Chat.ID,
		Text: `Assalomu alaykum!
Ushbu bot orqali siz Go dasturlash tiliga oid test savollaridan o'tishingiz mumkin.
Testni boshlash uchun quyidagi tugmani bosing!
👇👇👇👇👇👇👇👇👇`,
		ParseMode: "html",
		ReplyMarkup: &Keyboard{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "GO, 🚀",
						CallbackData: "/start_test",
					},
				},
			},
		},
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	s.logger.Debug().Any("response message", msg).Send()

	return nil
}
