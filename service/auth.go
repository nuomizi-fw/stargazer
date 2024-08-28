package service

type AuthService interface{}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}
