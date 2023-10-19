package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/handler/rest/user"
	"github.com/abdivasiyev/telegram-bot/pkg/config"
)

func New(p Params) {
	p.Router.Route("/user", func(r chi.Router) {
		r.Get("/", p.User.Create)
	})

	s := &http.Server{
		Handler: p.Router,
		Addr:    fmt.Sprintf(":%d", p.Config.GetInt("server.port")),
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(fx.Invoke(New))

type Params struct {
	fx.In
	Lifecycle fx.Lifecycle
	Config    config.Config
	Router    chi.Router
	User      user.Handler
}
