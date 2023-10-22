package user

import (
	"context"
	"errors"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
	"github.com/abdivasiyev/telegram-bot/internal/ent/store/user"
)

func (r *repo) Update(ctx context.Context, req UpdateUser) error {
	_, err := r.client.User.Update().
		SetQuizID(req.QuizID).
		SetIncorrectQuiz(req.Incorrects).
		SetCorrectQuiz(req.Corrects).
		Where(user.TelegramID(req.TelegramID)).
		Save(ctx)
	return err
}

func (r *repo) Get(ctx context.Context, telegramID int64) (*store.User, error) {
	u, err := r.client.User.Query().
		Where(user.TelegramID(telegramID)).First(ctx)

	if store.IsNotFound(err) {
		return nil, ErrUserNotFound
	}

	return u, nil
}

func (r *repo) Create(ctx context.Context, req CreateUser) (*store.User, error) {
	u, err := r.client.User.Create().
		SetUsername(req.Username).
		SetTelegramID(req.TelegramID).
		SetQuizID(1).
		Save(ctx)
	return u, err
}

var (
	Module = fx.Provide(New)

	ErrUserNotFound = errors.New("user not found")
)

func New(client *store.Client) Repo {
	return &repo{client: client}
}

type (
	UpdateUser struct {
		TelegramID int64
		QuizID     int64
		Incorrects int64
		Corrects   int64
	}

	CreateUser struct {
		Username   string
		TelegramID int64
	}

	repo struct {
		client *store.Client
	}

	Repo interface {
		Get(ctx context.Context, telegramID int64) (*store.User, error)
		Create(ctx context.Context, req CreateUser) (*store.User, error)
		Update(ctx context.Context, req UpdateUser) error
	}
)
