package bot

import (
	"net/http"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/service/bot"
	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

var Module = fx.Provide(New)

func New(p Params) Handler {
	return &handler{
		logger:  p.Logger,
		service: p.Service,
	}
}

type handler struct {
	service bot.Service
	logger  *zerolog.Logger
}

type Params struct {
	fx.In
	Logger  *zerolog.Logger
	Service bot.Service
}

type Handler interface {
	Webhook(w http.ResponseWriter, r *http.Request)
}
