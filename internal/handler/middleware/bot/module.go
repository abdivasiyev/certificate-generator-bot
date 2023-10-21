package bot

import (
	"io"
	"net/http"
	"strings"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/service/bot"
	"github.com/abdivasiyev/telegram-bot/pkg/config"
	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

func (h *handler) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			requestToken = r.Header.Get(telegramSecretTokenHeader)
			secretToken  = h.config.GetString("webhook.secret")
		)

		if strings.Compare(requestToken, secretToken) != 0 {
			w.WriteHeader(http.StatusOK)
			_, _ = io.WriteString(w, "Nice try :)")
			return
		}

		next.ServeHTTP(w, r)
	})
}

const telegramSecretTokenHeader = "X-Telegram-Bot-Api-Secret-Token"

var Module = fx.Provide(New)

func New(p Params) Handler {
	return &handler{
		logger:  p.Logger,
		config:  p.Config,
		service: p.Service,
	}
}

type (
	Params struct {
		fx.In
		Logger  *zerolog.Logger
		Config  config.Config
		Service bot.Service
	}

	handler struct {
		logger  *zerolog.Logger
		config  config.Config
		service bot.Service
	}

	Handler interface {
		Validate(next http.Handler) http.Handler
	}
)
