package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	EMAIL_SVC_PORT string `mapstructure:"EMAIL_SVC_PORT"`
	SMTP_HOST      string `mapstructure:"SMTP_HOST"`
	SMTP_PORT      int16  `mapstructure:"SMTP_PORT"`
	SMTP_USER      string `mapstructure:"SMTP_USER"`
	SMTP_PWD       string `mapstructure:"SMTP_PWD"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	// 	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	// 	TODO: Set config for permission type possibly in env

	return config, err
}
