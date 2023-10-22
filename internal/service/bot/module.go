package bot

import (
	"context"
	"errors"

	"go.uber.org/fx"

	"github.com/imroc/req/v3"

	"github.com/abdivasiyev/telegram-bot/internal/repo/quiz"
	"github.com/abdivasiyev/telegram-bot/internal/repo/user"
	"github.com/abdivasiyev/telegram-bot/pkg/config"
	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

var (
	ErrTelegramRequest = errors.New("error sending request to telegram")
	ErrDeleteWebhook   = errors.New("error deleting webhook")
	ErrSetWebhook      = errors.New("error setting webhook")

	Module = fx.Provide(New)
)

func New(p Params) Service {
	s := &service{
		logger: p.Logger,
		httpClient: req.NewClient().
			SetCommonPathParam("token", p.Config.GetString("bot.token")).
			SetBaseURL("https://api.telegram.org").
			SetCommonHeader("Content-Type", "application/json"),
		webhookURL:    p.Config.GetString("webhook.url"),
		webhookSecret: p.Config.GetString("webhook.secret"),
		botToken:      p.Config.GetString("bot.token"),
		userRepo:      p.UserRepo,
		quizRepo:      p.QuizRepo,
	}

	p.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return s.setWebhook(ctx)
		},
	})

	return s
}

type (
	Params struct {
		fx.In
		Lifecycle fx.Lifecycle
		Config    config.Config
		Logger    *zerolog.Logger
		UserRepo  user.Repo
		QuizRepo  quiz.Repo
	}

	service struct {
		logger        *zerolog.Logger
		webhookURL    string
		webhookSecret string
		botToken      string
		httpClient    *req.Client
		userRepo      user.Repo
		quizRepo      quiz.Repo
	}

	Service interface {
		Handle(ctx context.Context, update Update) error
	}
)
