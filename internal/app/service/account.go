package service

import (
	"errors"
	"github.com/Allswbr/balance-service/internal/app/repository"
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
