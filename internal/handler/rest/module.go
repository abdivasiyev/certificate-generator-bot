package rest

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/handler/rest/bot"
)

var Module = fx.Options(
	bot.Module,
)
