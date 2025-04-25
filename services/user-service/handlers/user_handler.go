package handlers

import "hauparty/services/user-service/services"

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}
