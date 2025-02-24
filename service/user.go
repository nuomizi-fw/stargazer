package service

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strings"
	"sync"
	"unicode"

	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/core"
	"github.com/nuomizi-fw/stargazer/pkg/jwt"
	"github.com/nuomizi-fw/stargazer/pkg/keystore"
	"github.com/nuomizi-fw/stargazer/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrWeakPassword       = errors.New("password too weak: must be at least 8 characters long and contain uppercase, lowercase, number, and special character")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidUsername    = errors.New("invalid username: must be 3-32 characters long and contain only letters, numbers, dots, and underscores")
)

type UserService interface {
	Register(register api.RegisterRequest) (api.User, error)
	Login(login api.LoginRequest) (api.Token, error)
	RefreshToken(refresh api.RefreshTokenRequest) (api.Token, error)

	GetUser() error
	GetUsers() error
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
	SetUserRole() error
	ResetPassword() error
}

var UserRegisterHash = sync.Map{}

type userService struct {
	logger     core.StargazerLogger
	repo       repository.Repository
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func validatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func validateEmail(email string) bool {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) == 0 || len(parts[1]) == 0 {
		return false
	}
	return true
}

func validateUsername(username string) bool {
	if len(username) < 3 || len(username) > 32 {
		return false
	}
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '.' && char != '_' {
			return false
		}
	}
	return true
}

func (u *userService) Register(register api.RegisterRequest) (api.User, error) {
	// Validate username format
	if !validateUsername(register.Username) {
		return api.User{}, ErrInvalidUsername
	}

	// Validate email format
	if !validateEmail(string(register.Email)) {
		return api.User{}, ErrInvalidEmail
	}

	// Validate password strength
	if !validatePassword(register.Password) {
		return api.User{}, ErrWeakPassword
	}

	// Check if user exists
	exists, err := u.repo.UserExists(register.Username)
	if err != nil {
		return api.User{}, fmt.Errorf("error checking user existence: %w", err)
	}
	if exists {
		return api.User{}, ErrUserExists
	}

	// Hash password with appropriate cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return api.User{}, fmt.Errorf("error hashing password: %w", err)
	}

	// Create user
	user := api.User{
		Username: register.Username,
		Email:    register.Email,
	}

	err = u.repo.CreateUser(user, string(hashedPassword))
	if err != nil {
		return api.User{}, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (u *userService) Login(login api.LoginRequest) (api.Token, error) {
	// Get user and password
	user, hashedPassword, err := u.repo.GetUserWithPassword(login.Username)
	if err != nil {
		// Don't expose internal errors to client
		return api.Token{}, ErrInvalidCredentials
	}

	// Verify password with constant-time comparison
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(login.Password))
	if err != nil {
		return api.Token{}, ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, err := jwt.GetAccessToken(user.Username, u.privateKey, user.Id, jwt.AccessTokenDuration)
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken, err := jwt.GetRefreshToken(user.Username, u.privateKey, user.Id, jwt.RefreshTokenDuration)
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	return api.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(jwt.AccessTokenDuration.Seconds()),
	}, nil
}

func (u *userService) RefreshToken(refresh api.RefreshTokenRequest) (api.Token, error) {
	// Validate refresh token and ensure it's not expired
	valid, claims, err := jwt.Validate(refresh.RefreshToken, func() (*ecdsa.PublicKey, error) {
		return u.publicKey, nil
	})
	if err != nil || !valid {
		return api.Token{}, ErrInvalidToken
	}

	// Generate new tokens
	accessToken, err := jwt.GetAccessToken(claims.Username, u.privateKey, claims.ID, jwt.AccessTokenDuration)
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating access token: %w", err)
	}

	refreshToken, err := jwt.GetRefreshToken(claims.Username, u.privateKey, claims.ID, jwt.RefreshTokenDuration)
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	return api.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(jwt.AccessTokenDuration.Seconds()),
	}, nil
}

func (u *userService) GetUser() error {
	return nil
}

func (u *userService) GetUsers() error {
	return nil
}

func (u *userService) CreateUser() error {
	return nil
}

func (u *userService) UpdateUser() error {
	return nil
}

func (u *userService) DeleteUser() error {
	return nil
}

func (u *userService) SetUserRole() error {
	return nil
}

func (u *userService) ResetPassword() error {
	return nil
}

func NewUserService(
	logger core.StargazerLogger,
	repo repository.Repository,
	ks *keystore.KeyStore,
) UserService {
	return &userService{
		logger: logger,
		repo:   repo,
	}
}
