package service

import (
	"github.com/Allswbr/balance-service/internal/app/repository"
	"github.com/Allswbr/balance-service/model"
)

// UserService - структура, методы которой реализуют логику для работы с пользователями
type UserService struct {
	repo repository.Users
}

// CreateUser создает пользователя и возвращает либо его id, либо 0 и ошибку
func (us *UserService) CreateUser(user *model.Users) (int64, error) {
	return us.repo.CreateUser(user)
}

// GetAllUsers позволяет посмотреть всех существующих в системе пользователей
// Возвращает либо список пользователей, либо nil и ошибку
func (us *UserService) GetAllUsers() ([]*model.Users, error) {
	return us.repo.GetAllUsers()
}
