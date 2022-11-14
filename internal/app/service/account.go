package service

import (
	"errors"
	"github.com/Allswbr/balance-service/internal/app/repository"
	"github.com/Allswbr/balance-service/model"
)

// AccountService - структура, методы которой реализуют логику для работы с банковским счетом
type AccountService struct {
	repo repository.Account
}

// NewAccountService - конструктор для BankAccountService
func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

// GetBalanceByUserID возвращает баланс пользователя с ID, равным userID
func (b *AccountService) GetBalanceByUserID(userID int64) (map[string]string, error) {
	return b.repo.GetBalanceByUserID(userID)
}

// DepositMoneyToUser начисляет amount денег пользователяю с userID
func (b *AccountService) DepositMoneyToUser(userID int64, amount float64, details string) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	return b.repo.DepositMoneyToUser(userID, amount, details)
}

// ReservationUserAccount резервирует amount денег пользователяю с userID за услугу с serveceID и заказ с serviceID
func (b *AccountService) ReservationUserAccount(userID int64, serviceID int64, orderID int64, amount float64, details string) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	return b.repo.ReservationUserAccount(userID, serviceID, orderID, amount, details)
}

// ConfessionOrder снимает amount денег с резервного счета пользователяю с userID за услугу с serveceID и заказ с serviceID
func (b *AccountService) ConfessionOrder(userID int64, serviceID int64, orderID int64, amount float64, details string) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	return b.repo.ConfessionOrder(userID, serviceID, orderID, amount, details)
}

// TransferMoneyBetweenUsers переводит amount денег со счета пользователя с fromUserID пользователю с toUserID
func (b *AccountService) TransferMoneyBetweenUsers(fromUserID int64, toUserID int64, amount float64, details string) error {
	if amount <= 0 {
		return errors.New("the amount must be greater than zero")
	}
	return b.repo.TransferMoneyBetweenUsers(fromUserID, toUserID, amount, details)
}

// GetTransactionsHistory возвращает историю транзакций пользователя с userID
func (b *AccountService) GetTransactionsHistory(userID int64) ([]*model.Transactions, error) {
	return b.repo.GetTransactionsHistory(userID)
}
