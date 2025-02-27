package handler

import (
	"firstGoProject/internal/domain/service"
)

type UserHandler struct {
	s service.UserServicePort
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s}
}
