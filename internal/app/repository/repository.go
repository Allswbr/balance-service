package repository

import "github.com/Allswbr/balance-service/model"

// Users - интерфейс для списка методов с пользователями в слое репозитория
type Users interface {
	CreateUser(*model.Users) (int64, error)
	GetAllUsers() ([]*model.Users, error)
}

// Account - интерфейс для списка методов с банковским счетом пользователя в слое репозитория
type Account interface {
}

// Repository — отвечает за получение данных из внешних источников, такие как база данных, api, локальное хранилище и пр.
type Repository struct {
	Users
	Account
}
