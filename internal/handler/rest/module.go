package rest

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/handler/rest/user"
)

var Module = fx.Options(
	user.Module,
)
