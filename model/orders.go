package model

import "time"

// Orders - структура для хранения проведённых заказов
type Orders struct {
	OrderID   int64     `json:"order_id"`
	UserID    int64     `json:"user_id"`
	ServiceID int64     `json:"service_id"`
	Amount    int64     `json:"amount"`
	Datetime  time.Time `json:"datetime"`
}
