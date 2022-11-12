package service

import "github.com/Allswbr/balance-service/model"

// Users - интерфейс для списка методов с пользователями в слое сервиса
type Users interface {
	CreateUser(*model.User) (int64, error)
	GetAllUsers() ([]*model.User, error)
	GetUserByID(userID int64) (*model.User, error)
}

// Account - интерфейс для списка методов с банковским счетом пользователя в слое сервиса
type Account interface {
}

// Service — отвечает за бизнес логику и ее переиспользование между компонентами
type Service struct {
	Users
	Account
}
