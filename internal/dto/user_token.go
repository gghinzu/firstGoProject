package dto

type TokenUserDTO struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
