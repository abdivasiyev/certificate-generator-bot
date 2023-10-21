package user

import (
	"context"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
	"github.com/abdivasiyev/telegram-bot/internal/ent/store/user"
)

func (r *repo) Update(ctx context.Context, req UpdateUser) error {
	_, err := r.client.User.Update().
		SetQuizID(req.QuizID).
		Where(user.TelegramID(req.TelegramID)).
		Save(ctx)
	return err
}

func (r *repo) Get(ctx context.Context, telegramID int64) (*store.User, error) {
	return r.client.User.Query().
		Where(user.TelegramID(telegramID)).First(ctx)
}

func (r *repo) Create(ctx context.Context, req CreateUser) error {
	_, err := r.client.User.Create().
		SetUsername(req.Username).
		SetTelegramID(req.TelegramID).
		Save(ctx)
	return err
}

var Module = fx.Provide(New)

func New(client *store.Client) Repo {
	return &repo{client: client}
}

type (
	UpdateUser struct {
		TelegramID int64
		QuizID     int64
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
	}
)
