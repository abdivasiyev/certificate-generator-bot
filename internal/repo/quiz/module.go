package quiz

import (
	"context"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
)

func (r *repo) Get(ctx context.Context, id int) (*store.Quiz, error) {
	return r.client.Quiz.Get(ctx, id)
}

var Module = fx.Provide(New)

func New(client *store.Client) Repo {
	return &repo{client: client}
}

type repo struct {
	client *store.Client
}

type Repo interface {
	Get(ctx context.Context, id int) (*store.Quiz, error)
}
