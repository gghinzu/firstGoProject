package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

// for user creation, DTO
type CreateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

// for updating user, DTO
type UpdateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Role      entity.UserRole `json:"role"`
}

// should bind (form) usage
type SearchUserDTO struct {
	Name      *string          `json:"name" form:"name"`
	Surname   *string          `json:"surname" form:"surname"`
	Age       *int             `json:"age" form:"age"`
	Gender    *enum.UserGender `json:"gender" form:"gender"`
	Education *string          `json:"education" form:"education"`
	Status    *enum.UserStatus `json:"status" form:"status"`
	Page      *int             `json:"page" form:"page"`
	Limit     *int             `json:"limit" form:"limit"`
}

type SignUpDTO struct {
	Email     string          `validate:"required" json:"email" gorm:"unique;not null"`
	Password  string          `validate:"required" json:"password" gorm:"not null"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

type LoginDTO struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type TokenUserDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenDTO struct {
	Token string `json:"token"`
}

type ConfigDTO struct {
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
