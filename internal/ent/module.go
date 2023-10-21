package ent

import (
	"context"

	"entgo.io/ent/dialect"
	"go.uber.org/fx"

	_ "github.com/lib/pq"

	"github.com/abdivasiyev/telegram-bot/internal/ent/store"
	"github.com/abdivasiyev/telegram-bot/pkg/config"
	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

var Module = fx.Invoke(New)

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.Config
	Logger    *zerolog.Logger
}

func New(p Params) (*store.Client, error) {
	url := p.Config.GetString("postgres.url")

	client, err := store.Open(dialect.Postgres, url)
	if err != nil {
		return nil, err
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			p.Logger.Info().Msg("running schema migration on startup")
			if err = client.Schema.Create(ctx); err != nil {
				p.Logger.Error().Err(err).Msg("failed to create schema")
				return err
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
