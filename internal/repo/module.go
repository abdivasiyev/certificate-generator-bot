package repo

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/repo/quiz"
	"github.com/abdivasiyev/telegram-bot/internal/repo/user"
)

var Module = fx.Options(
	user.Module,
	quiz.Module,
)
