package service

import (
	"github.com/Allswbr/balance-service/internal/app/repository"
	"github.com/Allswbr/balance-service/model"
)

// User - интерфейс для списка методов с пользователями в слое сервиса
type User interface {
	CreateUser(user *model.User) (int64, error)
	GetAllUsers() ([]*model.User, error)
	GetUserByID(userID int64) (*model.User, error)
}

// Account - интерфейс для списка методов с банковским счетом пользователя в слое сервиса
type Account interface {
	GetBalanceByUserID(userID int64) (map[string]string, error)
	DepositMoneyToUser(userID int64, amount float64, details string) error
	ReservationUserAccount(userID int64, serviceID int64, orderID int64, amount float64, details string) error
	ConfessionOrder(userID int64, serviceID int64, orderID int64, amount float64, details string) error
	TransferMoneyBetweenUsers(fromUserID int64, toUserID int64, amount float64, details string) error
	GetTransactionsHistory(userID int64) ([]*model.Transactions, error)
}

// Service — отвечает за бизнес логику и ее переиспользование между компонентами
type Service struct {
	User
	Account
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    NewUserService(repo.User),
		Account: NewAccountService(repo),
	}
}
