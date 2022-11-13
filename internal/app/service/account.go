package service

import "github.com/Allswbr/balance-service/internal/app/repository"

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
