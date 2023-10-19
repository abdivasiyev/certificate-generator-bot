package router

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/pkg/router/chi"
)

var Module = fx.Provide(chi.New)
