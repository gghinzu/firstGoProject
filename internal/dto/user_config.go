package dto

type ConfigDTO struct {
	ClientOrigin     string `mapstructure:"CLIENT_ORIGIN"`
	JWTAccessSecret  string `mapstructure:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret string `mapstructure:"JWT_REFRESH_SECRET"`

	SMTPHost           string `mapstructure:"SMTP_HOST"`
	SMTPPort           string `mapstructure:"SMTP_PORT"`
	SMTPSenderEmail    string `mapstructure:"SMTP_SENDER_EMAIL"`
	SMTPSenderPassword string `mapstructure:"SMTP_SENDER_PASSWORD"`
}
