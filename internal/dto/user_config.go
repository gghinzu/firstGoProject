package dto

type ConfigDTO struct {
	ClientOrigin     string `mapstructure:"CLIENT_ORIGIN"`
	JWTAccessSecret  string `mapstructure:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`
}
