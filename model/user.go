package model

import "github.com/go-playground/validator/v10"

// User - структура для представления пользователей и их счетов
type User struct {
	UserID         int64 `json:"user_id" db:"user_id"`
	DepositAccount int64 `json:"deposit_account" db:"deposit_account"`
	ReserveAccount int64 `json:"reserve_account" db:"reserve_account"`
}

var validate *validator.Validate

// Validate проверяет валидность полей пользователя
func (u *User) Validate() error {
	validate = validator.New()
	return validate.Struct(u)
}
