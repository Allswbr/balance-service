package repository

import (
	"fmt"
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
)

// AccountPostgres - структура, которая отвечает за связь работы с балансом в PostgreSQL
type AccountPostgres struct {
	db *sqlx.DB
}

// NewAccountPostgres - конструктор для BankAccountPostgres
func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

// GetBalanceByUserID возвращает баланс пользователя с ID, равным userID
func (b *AccountPostgres) GetBalanceByUserID(userID int64) (map[string]string, error) {
	user := &model.User{}
	err := b.db.Get(&user, "SELECT * FROM users WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	userDeposit := fmt.Sprintf("%.2f", user.DepositAccount)
	userReserve := fmt.Sprintf("%.2f", user.ReserveAccount)
	return map[string]string{
		"deposit": userDeposit,
		"reserve": userReserve,
	}, nil
}
