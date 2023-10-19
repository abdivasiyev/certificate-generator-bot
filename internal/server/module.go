package server

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/server/http"
)

var Module = fx.Options(
	http.Module,
)
