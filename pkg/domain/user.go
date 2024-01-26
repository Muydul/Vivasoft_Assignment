package domain

import (
	"book-crud/pkg/models"
	"book-crud/pkg/types"
)

type IUserRepo interface {
	GetUser(username *string) (*models.UserDetail, error)
	CreateUser(user *models.UserDetail) error
}

type IAuthService interface {
	LoginUser(user *types.LoginRequest) (string, error)
	SignupUser(user *types.UserRequest) error
}