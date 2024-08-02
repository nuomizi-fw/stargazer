package service

import "gorm.io/gorm"

type UserService interface{}

type userService struct {
	db *gorm.DB
}

func NewUserService() UserService {
	return &userService{}
}
