package service

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/nuomizi-fw/stargazer/api"
	"github.com/nuomizi-fw/stargazer/pkg/jwt"
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

type AuthService interface {
	Register(register api.RegisterRequest) (api.User, error)
	Login(login api.LoginRequest) (api.Token, error)
	RefreshToken(refresh api.RefreshTokenRequest) (api.Token, error)
}

type authService struct {
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

func (a *authService) Register(register api.RegisterRequest) (api.User, error) {
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
	exists, err := a.repo.UserExists(register.Username)
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

	err = a.repo.CreateUser(user, string(hashedPassword))
	if err != nil {
		return api.User{}, fmt.Errorf("error creating user: %w", err)
	}

	return user, nil
}

func (a *authService) Login(login api.LoginRequest) (api.Token, error) {
	// Get user and password
	user, hashedPassword, err := a.repo.GetUserWithPassword(login.Username)
	if err != nil {
		// Don't expose internal errors to client
		return api.Token{}, ErrInvalidCredentials
	}

	// Verify password with constant-time comparison
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(login.Password))
	if err != nil {
		return api.Token{}, ErrInvalidCredentials
	}

	// Generate tokens based on client type
	var accessToken, refreshToken string
	var expiresIn int

	switch login.ClientType {
	case "web":
		accessToken, err = jwt.GetWebAccessToken(user.Username, a.privateKey, user.Id)
		if err != nil {
			return api.Token{}, fmt.Errorf("error generating web access token: %w", err)
		}
		refreshToken, err = jwt.GetWebRefreshToken(user.Username, a.privateKey, user.Id)
		expiresIn = int(jwt.WebAccessTokenDuration.Seconds())
	case "native":
		accessToken, err = jwt.GetNativeAccessToken(user.Username, a.privateKey, user.Id)
		if err != nil {
			return api.Token{}, fmt.Errorf("error generating native access token: %w", err)
		}
		refreshToken, err = jwt.GetNativeRefreshToken(user.Username, a.privateKey, user.Id)
		expiresIn = int(jwt.NativeAccessTokenDuration.Seconds())
	default:
		return api.Token{}, fmt.Errorf("invalid client type: %s", login.ClientType)
	}
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	return api.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

func (a *authService) RefreshToken(refresh api.RefreshTokenRequest) (api.Token, error) {
	// Validate refresh token and ensure it's not expired
	valid, claims, err := jwt.Validate(refresh.RefreshToken, func() (*ecdsa.PublicKey, error) {
		return a.publicKey, nil
	})
	if err != nil || !valid {
		return api.Token{}, ErrInvalidToken
	}

	// Ensure the token is a refresh token
	if claims.TokenType != jwt.TokenTypeRefresh {
		return api.Token{}, ErrInvalidToken
	}

	// Check if the token is about to expire and reject if too close to expiry
	if claims.ExpiresAt != nil && time.Until(claims.ExpiresAt.Time) < 5*time.Minute {
		return api.Token{}, ErrInvalidToken
	}

	// Generate new tokens based on client type
	var accessToken, refreshToken string
	var expiresIn int

	switch claims.ClientType {
	case string(api.Web):
		accessToken, err = jwt.GetWebAccessToken(claims.Username, a.privateKey, claims.ID)
		if err != nil {
			return api.Token{}, fmt.Errorf("error generating web access token: %w", err)
		}
		refreshToken, err = jwt.GetWebRefreshToken(claims.Username, a.privateKey, claims.ID)
		expiresIn = int(jwt.WebAccessTokenDuration.Seconds())
	case string(api.Native):
		accessToken, err = jwt.GetNativeAccessToken(claims.Username, a.privateKey, claims.ID)
		if err != nil {
			return api.Token{}, fmt.Errorf("error generating native access token: %w", err)
		}
		refreshToken, err = jwt.GetNativeRefreshToken(claims.Username, a.privateKey, claims.ID)
		expiresIn = int(jwt.NativeAccessTokenDuration.Seconds())
	default:
		return api.Token{}, ErrInvalidToken
	}
	if err != nil {
		return api.Token{}, fmt.Errorf("error generating refresh token: %w", err)
	}

	return api.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

func NewAuthService(
	repo repository.Repository,
	privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey,
) AuthService {
	return &authService{
		repo:       repo,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}
