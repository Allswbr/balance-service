package repository

import (
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
)

// User - интерфейс для списка методов с пользователями в слое репозитория
type User interface {
	CreateUser(user *model.User) (int64, error)
	GetAllUsers() ([]*model.User, error)
	GetUserByID(userID int64) (*model.User, error)
}

// Account - интерфейс для списка методов с банковским счетом пользователя в слое репозитория
type Account interface {
	GetBalanceByUserID(userID int64) (map[string]string, error)
	DepositMoneyToUser(userID int64, amount float64, details string) error
	ConfessionOrder(userID int64, serviceID int64, orderID int64, amount float64, details string) error
	ReservationUserAccount(userID int64, serviceID int64, orderID int64, amount float64, details string) error
	TransferMoneyBetweenUsers(fromUserID int64, toUserID int64, amount float64, details string) error
	GetTransactionsHistory(userID int64) ([]*model.Transactions, error)
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
		Account: NewAccountPostgres(db),
	}
}
