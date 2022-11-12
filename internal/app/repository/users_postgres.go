package repository

import (
	"fmt"
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

// UsersPostgres - структура, которая отвечает за свзяь работы с пользователями и PostgreSQL
type UsersPostgres struct {
	db *sqlx.DB
}

// CreateUser создает пользователя и возвращает либо его id, либо 0 и ошибку
func (up *UsersPostgres) CreateUser(user *model.Users) (int64, error) {
	err := up.db.QueryRow(
		`INSERT INTO users (deposit_account, reserve_account) VALUES ($1, $2) RETURNING user_id`,
		0, 0,
	).Scan(&user.UserID)

	if err != nil {
		if strings.Contains(err.Error(), "ER_DUP_ENTRY") {
			return 0, fmt.Errorf("failed to add a new user")
		} else {
			return 0, err
		}
	}

	return user.UserID, nil
}

// GetAllUsers позволяет посмотреть всех существующих в системе пользователей
// Возвращает либо список пользователей, либо nil и ошибку
func (up *UsersPostgres) GetAllUsers() ([]*model.Users, error) {
	rows, err := up.db.Queryx("SELECT (user_id, deposit_account, reserve_account) FROM users")
	if err != nil {
		return nil, err
	}

	users := make([]*model.Users, 0)
	u := &model.Users{}
	for rows.Next() {
		if err = rows.StructScan(u); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
