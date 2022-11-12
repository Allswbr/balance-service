package model

// Users - структура для представления пользователей и их счетов
type Users struct {
	UserID         int64 `json:"user_id"`
	DepositAccount int64 `json:"deposit_account"`
	ReserveAccount int64 `json:"reserve_account"`
}
