package model

import "time"

// Transactions - структура для сохранения денежных переводов
type Transactions struct {
	TransID      int64     `json:"trans_id"`
	UserID       int64     `json:"user_id"`
	ServiceID    int64     `json:"service_id"`
	OrderID      int64     `json:"order_id"`
	Datetime     time.Time `json:"datetime"`
	Amount       float64   `json:"amount"`
	StartDeposit float64   `json:"start_deposit"`
	EndDeposit   float64   `json:"end_deposit"`
	Event        string    `json:"event"`
	Message      string    `json:"message"`
}
