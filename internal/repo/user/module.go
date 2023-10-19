package user

import (
	"context"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
)

func (r *repo) Create(ctx context.Context) error {
	_, err := r.client.User.Create().
		SetUsername("abdivasiyev").
		SetTelegramID(1).
		Save(ctx)
	return err
}

var Module = fx.Provide(New)

func New(client *store.Client) Repo {
	return &repo{client: client}
}

type repo struct {
	client *store.Client
}

type Repo interface {
	Create(ctx context.Context) error
}
