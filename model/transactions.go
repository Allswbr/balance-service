package model

import "time"

// Transactions - структура для сохранения денежных переводов
type Transactions struct {
	TransID      int64     `json:"trans_id" db:"trans_id"`
	UserID       int64     `json:"user_id" db:"user_id"`
	ServiceID    int64     `json:"service_id" db:"service_id"`
	OrderID      int64     `json:"order_id" db:"order_id"`
	Datetime     time.Time `json:"datetime" db:"datetime"`
	Amount       float64   `json:"amount" db:"amount"`
	StartDeposit float64   `json:"start_deposit" db:"start_deposit"`
	EndDeposit   float64   `json:"end_deposit" db:"end_deposit"`
	Event        string    `json:"event" db:"event"`
	Message      string    `json:"message" db:"message"`
}
