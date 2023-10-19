package zerolog

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New() *Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	return &Logger{
		zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Logger(),
	}
}
