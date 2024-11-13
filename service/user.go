package service

import (
	"crypto/ecdsa"
	"sync"

	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/pkg/jwt"
	"github.com/nuomizi-fw/stargazer/pkg/logger"
	"github.com/nuomizi-fw/stargazer/repository"
)

type UserService interface {
	GetUser() error
	GetUsers() error
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
	SetUserRole() error
	ResetPassword() error
	RefreshToken() error
	GetKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey)
}

var UserRegisterHash = sync.Map{}

type userService struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey

	repo repository.Repository
}

func NewUserService(
	repo repository.Repository,
) UserService {
	privateKey, publicKey, err := jwt.GenerateKeyPair()
	if err != nil {
		logger.Errorf("Failed to generate key pair: %s", err)
		return nil
	}

	return &userService{
		privateKey: privateKey,
		publicKey:  publicKey,
		repo:       repo,
	}
}

func (us *userService) GetUser() error {
	return nil
}

func (us *userService) GetUsers() error {
	return nil
}

func (us *userService) CreateUser() error {
	return nil
}

func (us *userService) UpdateUser() error {
	return nil
}

func (us *userService) DeleteUser() error {
	return nil
}

func (us *userService) SetUserRole() error {
	return nil
}

func (us *userService) ResetPassword() error {
	return nil
}

func (ur *userService) RefreshToken() error {
	return nil
}

func (u *userService) GetKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	return u.privateKey, u.publicKey
}
