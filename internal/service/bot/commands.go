package bot

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/abdivasiyev/telegram-bot/internal/repo/user"
)

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

	var (
		startText = userGotTextOnStart
	)

	u, err := s.userRepo.Get(ctx, message.Chat.ID)
	if err != nil {
		s.logger.Error().Err(err).Msg("error while getting user from database")

		switch {
		case errors.Is(err, user.ErrUserNotFound):
			u, err = s.userRepo.Create(ctx, user.CreateUser{
				Username:   message.Chat.Username,
				TelegramID: message.Chat.ID,
			})

			if err != nil {
				s.logger.Error().Err(err).Msg("error while creating user")
				return err
			}

			startText = userCreatedTextOnStart
		default:
			return err
		}
	}

	startText = strings.ReplaceAll(startText, "%quiz_id%", fmt.Sprintf("%04d", u.QuizID))

	msg, err := s.sendMessage(ctx, SendMessage{
		ChatID:    message.Chat.ID,
		Text:      startText,
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
