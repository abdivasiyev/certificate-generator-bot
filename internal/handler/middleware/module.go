package middleware

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/handler/middleware/bot"
)

var Module = fx.Options(
	bot.Module,
)
