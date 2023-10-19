package pkg

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/pkg/config"
	"github.com/abdivasiyev/telegram-bot/pkg/logger"
	"github.com/abdivasiyev/telegram-bot/pkg/router"
)

var Module = fx.Options(
	config.Module,
	logger.Module,
	router.Module,
)
