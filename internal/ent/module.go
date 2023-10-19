package ent

import (
	"context"

	"entgo.io/ent/dialect"
	"go.uber.org/fx"

	_ "github.com/lib/pq"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
	"github.com/abdivasiyev/telegram-bot/pkg/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.Config
}

func New(p Params) (*store.Client, error) {
	url := p.Config.GetString("postgres.url")

	client, err := store.Open(dialect.Postgres, url)
	if err != nil {
		return nil, err
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Schema.Create(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
