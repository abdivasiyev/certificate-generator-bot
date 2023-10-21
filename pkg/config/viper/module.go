package viper

import (
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type Viper struct {
	*viper.Viper
}

func New() *Viper {
	v := &Viper{viper.New()}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath(rootDir() + "/config")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	v.AutomaticEnv()
	v.WatchConfig()

	return v
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(path.Dir(path.Dir(b))))
	return filepath.Dir(d)
}
