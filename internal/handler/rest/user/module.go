package user

import (
	"io"
	"net/http"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/repo/user"
	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	if err := h.userRepo.Create(r.Context()); err != nil {
		h.logger.Debug().Err(err).Send()
	}

	_, _ = io.WriteString(w, "User create route")
}

var Module = fx.Provide(New)

func New(p Params) Handler {
	return &handler{
		userRepo: p.UserRepo,
		logger:   p.Logger,
	}
}

type handler struct {
	userRepo user.Repo
	logger   *zerolog.Logger
}

type Params struct {
	fx.In
	UserRepo user.Repo
	Logger   *zerolog.Logger
}

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
}
