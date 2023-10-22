package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
	"github.com/abdivasiyev/telegram-bot/internal/repo/user"
)

func (s *service) handleTestOption(ctx context.Context, callback *CallbackQuery) error {
	var currentQuizID = 1

	u, err := s.userRepo.Get(ctx, callback.Message.Chat.ID)
	if err != nil {
		s.logger.Error().Err(err).Msg("error while getting user from database")

		return fmt.Errorf("handleStartTest: %w", err)
	}

	q, err := s.quizRepo.Get(ctx, int(u.QuizID))
	if err != nil {
		s.logger.Error().Err(err).Msg("error while getting quiz from database")
		return fmt.Errorf("handleStartTest: %w", err)
	}

	// TODO: check question is correct or not
	if strings.TrimLeft(strings.TrimSpace(callback.Data), "/option_") == q.CorrectOption {
		u.CorrectQuiz++
	} else {
		u.IncorrectQuiz++
	}

	if q.NextQuizID > 0 {
		nextQ, err := s.quizRepo.Get(ctx, int(q.NextQuizID))
		if err != nil {
			s.logger.Error().Err(err).Msg("error while getting quiz from database")
			return fmt.Errorf("handleStartTest: %w", err)
		}

		if err = s.sendQuiz(ctx, nextQ, callback.Message.Chat.ID); err != nil {
			s.logger.Error().Err(err).Msg("failed to send next quiz")
			return err
		}

		currentQuizID = nextQ.ID
	} else {
		text := strings.ReplaceAll(quizEndText, "%corrects%", fmt.Sprintf("%02d", u.CorrectQuiz))
		text = strings.ReplaceAll(text, "%incorrects%", fmt.Sprintf("%02d", u.IncorrectQuiz))

		_, err = s.sendMessage(ctx, SendMessage{
			ChatID:    callback.Message.Chat.ID,
			Text:      text,
			ParseMode: "html",
			ReplyMarkup: &Keyboard{
				InlineKeyboard: [][]InlineKeyboardButton{
					{
						{
							Text:         "Sertifikatni qabul qilib olish",
							CallbackData: "/generate_certificate",
						},
					},
				},
			},
		})

		if err != nil {
			s.logger.Error().Err(err).Msg("failed to send message")
			return err
		}
	}

	if err = s.userRepo.Update(ctx, user.UpdateUser{
		TelegramID: callback.Message.Chat.ID,
		QuizID:     int64(currentQuizID),
		Corrects:   u.CorrectQuiz,
		Incorrects: u.IncorrectQuiz,
	}); err != nil {
		s.logger.Error().Err(err).Msg("failed to update user current quiz")
		return err
	}

	return nil
}

func (s *service) handleStartTest(ctx context.Context, callback *CallbackQuery) error {
	u, err := s.userRepo.Get(ctx, callback.Message.Chat.ID)
	if err != nil {
		s.logger.Error().Err(err).Msg("error while getting user from database")

		return fmt.Errorf("handleStartTest: %w", err)
	}

	q, err := s.quizRepo.Get(ctx, int(u.QuizID))
	if err != nil {
		s.logger.Error().Err(err).Msg("error while getting quiz from database")
		return fmt.Errorf("handleStartTest: %w", err)
	}

	_, err = s.editMessage(ctx, EditMessage{
		ChatID:    callback.Message.Chat.ID,
		MessageID: callback.Message.MessageID,
		Text:      "Test boshlandi!",
		ParseMode: "html",
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	if err = s.sendQuiz(ctx, q, callback.Message.Chat.ID); err != nil {
		s.logger.Error().Err(err).Msg("failed to send quiz")
		return err
	}

	return nil
}

func (s *service) sendQuiz(ctx context.Context, q *store.Quiz, chatID int64) error {
	quizMessage := strings.ReplaceAll(quizText, "%quiz_id%", fmt.Sprintf("%04d", q.ID))
	quizMessage = strings.ReplaceAll(quizMessage, "%quiz_question%", q.Question)
	quizMessage = strings.ReplaceAll(quizMessage, "%option_a%", q.OptionA)
	quizMessage = strings.ReplaceAll(quizMessage, "%option_b%", q.OptionB)
	quizMessage = strings.ReplaceAll(quizMessage, "%option_c%", q.OptionC)
	quizMessage = strings.ReplaceAll(quizMessage, "%option_d%", q.OptionD)

	_, err := s.sendMessage(ctx, SendMessage{
		ChatID:    chatID,
		Text:      quizMessage,
		ParseMode: "html",
		ReplyMarkup: &Keyboard{
			InlineKeyboard: [][]InlineKeyboardButton{
				{
					{
						Text:         "A",
						CallbackData: "/option_a",
					},
					{
						Text:         "B",
						CallbackData: "/option_b",
					},
				},
				{
					{
						Text:         "C",
						CallbackData: "/option_c",
					},
					{
						Text:         "D",
						CallbackData: "/option_d",
					},
				},
			},
		},
	})

	if err != nil {
		s.logger.Error().Err(err).Msg("failed to send message")
		return err
	}

	return nil
}
