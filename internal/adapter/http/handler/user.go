package handler

import "github.com/alihdrii/go-hexagonal/internal/core/usecase/user"

type UserHandler struct {
	UserUsecase *user.Usecase
}

func NewUserHandler(uc *user.Usecase) *UserHandler {
	return &UserHandler{UserUsecase: uc}
}
