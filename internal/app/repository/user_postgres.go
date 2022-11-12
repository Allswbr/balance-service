package repository

import (
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
)

// UserPostgres - структура, которая отвечает за свзяь работы с пользователями и PostgreSQL
type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres - конструктор для UserPostgres
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// CreateUser создает пользователя и возвращает либо его id, либо 0 и ошибку
func (up *UserPostgres) CreateUser(user *model.User) (int64, error) {
	err := up.db.QueryRow(
		`INSERT INTO users (deposit_account, reserve_account) VALUES ($1, $2) RETURNING user_id`,
		0, 0,
	).Scan(&user.UserID)

	if err != nil {
		return 0, err
	}

	return user.UserID, nil
}

// GetAllUsers позволяет посмотреть всех существующих в системе пользователей
// Возвращает либо список пользователей, либо nil и ошибку
func (up *UserPostgres) GetAllUsers() ([]*model.User, error) {

	users := make([]*model.User, 0)
	err := up.db.Select(&users, "SELECT (user_id, deposit_account, reserve_account) FROM users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID возвращает либо пользователя с id == userID
// либо nil и ошибку
func (up *UserPostgres) GetUserByID(userID int64) (*model.User, error) {

	user := &model.User{}
	if err := up.db.Get(user, "SELECT * FROM users WHERE id=$1", userID); err != nil {
		return nil, err
	}

	return user, nil
}
