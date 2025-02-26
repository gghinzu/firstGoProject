package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB_NAME"`

	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBSource     string `mapstructure:"DB_SOURCE"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	SeedMail     string `mapstructure:"PGADMIN_SEED_EMAIL"`
	SeedPassword string `mapstructure:"PGADMIN_SEED_PASSWORD"`
	SeedName     string `mapstructure:"PGADMIN_SEED_NAME"`
	SeedSurname  string `mapstructure:"PGADMIN_SEED_SURNAME"`

	JWTSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
