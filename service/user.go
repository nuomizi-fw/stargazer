package service

type UserService interface {
	GetUser() error
	GetUsers() error
	CreateUser() error
	UpdateUser() error
	DeleteUser() error
	SetUserRole() error
	SetUserPermission() error
	ResetPassword() error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
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

func (us *userService) SetUserPermission() error {
	return nil
}

func (us *userService) ResetPassword() error {
	return nil
}
