package handler

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/handler/rest"
)

var Module = fx.Options(rest.Module)
