package config

import (
	"firstGoProject/internal/dto"
	"github.com/spf13/viper"
)

func LoadConfig() (config dto.ConfigDTO, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	//TODO: it ignores .env's
	//bind env
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
