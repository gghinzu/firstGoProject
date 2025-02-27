package dto

type TokenUserDTO struct {
	Token        string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}
