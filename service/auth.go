package service

type AuthService interface {
	Register() error
	Login() error
	MFALogin() error
	ForgotPassword() error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (as *authService) Register() error {
	return nil
}

func (as *authService) Login() error {
	return nil
}

func (as *authService) MFALogin() error {
	return nil
}

func (as *authService) ForgotPassword() error {
	return nil
}
