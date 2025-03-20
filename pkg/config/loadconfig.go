package config

import (
	"firstGoProject/internal/dto"
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() (config dto.ConfigDTO, err error) {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("no .env file found: %v", err)
	} else {
		log.Println("successfully loaded .env file")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	if config == (dto.ConfigDTO{}) {
		log.Println("config is empty!")
	}

	return config, nil
}
