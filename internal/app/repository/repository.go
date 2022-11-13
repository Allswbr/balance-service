package repository

import (
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
)

// User - интерфейс для списка методов с пользователями в слое репозитория
type User interface {
	CreateUser(*model.User) (int64, error)
	GetAllUsers() ([]*model.User, error)
	GetUserByID(userID int64) (*model.User, error)
}

// Account - интерфейс для списка методов с банковским счетом пользователя в слое репозитория
type Account interface {
}

// Repository — отвечает за получение данных из внешних источников, такие как база данных, api, локальное хранилище и пр.
type Repository struct {
	User
	Account
}

// NewRepository - конструктор для Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Account: nil,
	}
}
