package services

import (
	"book-crud/pkg/domain"
	"book-crud/pkg/models"
	"book-crud/pkg/types"
	"book-crud/pkg/utils"
	"errors"
)

// authService defines the methods of the domain.IAuthService interface.
type authService struct {
	userRepo domain.IUserRepo
}

// AuthServiceInstance returns a new instance of the authService struct.
func AuthServiceInstance(userRepo domain.IUserRepo) domain.IAuthService {
	return &authService{
		userRepo: userRepo,
	}
}

// LoginUser returns a JWT token for the user if the credentials are correct.
func (service *authService) LoginUser(user *types.LoginRequest) (string, error) {
	// Check if user exists
	existingUser, err := service.userRepo.GetUser(&user.UserName)
	if err != nil {
		return "", errors.New("user does not exist")
	}

	// Check if password is correct
	if err := utils.CheckPassword(existingUser.PasswordHash, user.Password); err != nil {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token
	token, err := utils.GetJwtForUser(existingUser.Username)
	if err != nil {
		return "", err
	}

	return token, nil

}

// SignupUser creates a new user with the given user details.
func (service *authService) SignupUser(userRequest *types.UserRequest) error {
	// get hashed password
	passwordHash, err := utils.GetPasswordHash(userRequest.Password)
	if err != nil {
		return err
	}

	// create user
	user := &models.UserDetail{
		Username:     userRequest.Username,
		PasswordHash: passwordHash,
		Name:         userRequest.Name,
		Email:        userRequest.Email,
		Address:      userRequest.Address,
	}
	if err := service.userRepo.CreateUser(user); err != nil {
		return err
	}

	return nil
}