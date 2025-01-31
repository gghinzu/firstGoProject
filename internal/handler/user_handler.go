package handler

import (
	"firstGoProject/internal/domain/service"
)

// UserHandler is a structure to provide a communication with other layers
// we use structures for this purpose because we don't have 'implements' keyword in Go, for interface applications
// it communicates with the service
type UserHandler struct {
	service service.UserServicePort
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
