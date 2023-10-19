package main

import (
	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/internal"
	"github.com/abdivasiyev/telegram-bot/pkg"
)

func main() {
	fx.New(
		pkg.Module,
		internal.Module,
	).Run()
}
