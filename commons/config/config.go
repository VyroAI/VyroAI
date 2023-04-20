package config

import (
	"github.com/spf13/viper"
	"os"
)

func ReadConfig[T any]() (*T, error) {
	viper.AddConfigPath(".")

	if os.Getenv("LIGHTSTEP_ENV") == "prod" || os.Getenv("ENV") == "prod" {
		viper.SetConfigName("config.prod")
	} else {
		viper.SetConfigName("config.dev")
	}

	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg T
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
