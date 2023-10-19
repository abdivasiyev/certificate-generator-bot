package config

import (
	"time"

	"go.uber.org/fx"

	"github.com/abdivasiyev/telegram-bot/pkg/config/viper"
)

type Config interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetFloat64(key string) float64
	GetInt(key string) int
	GetUInt32(key string) uint32
	GetUInt8(key string) uint8
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	UnmarshalKey(key string, val interface{}) error
	GetStringSlice(key string) []string
	GetDuration(key string) time.Duration
}

var Module = fx.Provide(
	fx.Annotate(viper.New, fx.As(new(Config))),
)
