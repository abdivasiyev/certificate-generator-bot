package logger

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/pkg/logger/zerolog"
)

var Module = fx.Provide(zerolog.New)
