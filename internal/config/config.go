package config

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
)

// Init 初始化viper
func New() (*viper.Viper, error) {
	var (
		err error
		v   = viper.New()
	)

	v.SetEnvPrefix("IOJ")

	err = v.BindEnv("host")
	if err != nil {
		return nil, err
	}

	v.SetDefault("host", "http://10.20.107.171:2333")

	return v, err
}

var ProviderSet = wire.NewSet(New)
