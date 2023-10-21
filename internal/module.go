package internal

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal/ent"
	"github.com/abdivasiyev/telegram-bot/internal/handler"
	"github.com/abdivasiyev/telegram-bot/internal/repo"
	"github.com/abdivasiyev/telegram-bot/internal/server"
	"github.com/abdivasiyev/telegram-bot/internal/service"
)

var Module = fx.Options(
	ent.Module,
	repo.Module,
	handler.Module,
	server.Module,
	service.Module,
)
