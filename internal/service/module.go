package service

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/service/bot"
)

var Module = fx.Options(
	bot.Module,
)
