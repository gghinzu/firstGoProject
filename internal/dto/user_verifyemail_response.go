package dto

type VerifyEmailDTO struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}
