package repository

import (
	"errors"
	"fmt"
	"github.com/Allswbr/balance-service/model"
	"github.com/jmoiron/sqlx"
	"time"
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
	err := b.db.Get(user, "SELECT * FROM users WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}

	//userDeposit := fmt.Sprintf("%.2f", user.DepositAccount)
	//userReserve := fmt.Sprintf("%.2f", user.ReserveAccount)
	fmt.Println(user.DepositAccount)
	fmt.Println(user.ReserveAccount)
	return map[string]string{
		"deposit": fmt.Sprintf("%.f", user.DepositAccount),
		"reserve": fmt.Sprintf("%.f", user.ReserveAccount),
	}, nil
}

// DepositMoneyToUser начисляет amount денег пользователяю с userID
func (b *AccountPostgres) DepositMoneyToUser(userID int64, amount float64, details string) error {

	// Сразу сформируем сообщение для записи в таблицу транзакций
	msg := fmt.Sprintf("deposit %.2f to account", amount)
	if len(details) != 0 {
		msg += ": " + details
	}

	account := &model.User{}
	err := b.db.Get(account, "SELECT * FROM users WHERE user_id=$1", userID)

	// Кейс : у пользователя уже есть запись в таблице
	// значит надо обновить это значение
	if err != nil {
		return b.deposit(userID, amount, msg, account.DepositAccount)
	}
	return err
}

// deposit реализует взнос денег на счет пользователя
func (b *AccountPostgres) deposit(userID int64, amount float64, msg string, startDeposit float64) error {
	// Надо сделать 2 вещи: обновить баланс пользователя
	// И добавить соответствующую запись в таблицу историй транзакций
	// Объединим эти операции в транзакцию
	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	endDeposit := startDeposit + amount
	_, err = tx.Exec(
		"UPDATE users SET deposit_account=$1 WHERE user_id=$2",
		endDeposit,
		userID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Отрицательное значение передается при снятии денег
	// Так как в таблице есть ограничение на то, что сумма только положительная
	// Умножим ее на минус 1
	if amount < 0 {
		amount *= -1
	}
	_, err = tx.Exec(
		"INSERT INTO transaction (user_id, datetime, amount, start_deposit, end_deposit, event, message)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userID, time.Now(), amount, startDeposit, endDeposit, "ADD", msg,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// ReservationUserAccount резервирует amount денег пользователя с userID за услугу с serveceID и заказ с serviceID
func (b *AccountPostgres) ReservationUserAccount(userID int64, serviceID int64, orderID int64, amount float64, details string) error {

	account := &model.User{}
	err := b.db.Get(account, "SELECT * FROM users WHERE user_id=$1", userID)

	// Проверим, достаточно ли у пользователя денег
	if amount > account.DepositAccount {
		return errors.New("not enough money in the bank account")
	}

	// Cформируем сообщение для записи в таблицу транзакций
	msg := fmt.Sprintf("reserve %.2f on account", amount)
	if len(details) != 0 {
		msg += ": " + details
	}

	// Кейс : у пользователя уже есть запись в таблице
	// значит надо обновить это значение
	if err != nil {
		return b.reserve(userID, amount, msg, serviceID, orderID, account.DepositAccount, account.ReserveAccount)
	}
	return err
}

// reserve резервирует amount денег пользователя с userID за услугу с serveceID и заказ с serviceID
func (b *AccountPostgres) reserve(userID int64, amount float64, msg string, serviceID int64, orderID int64, startDeposit float64, startReserve float64) error {
	// Надо сделать 2 вещи: обновить счета пользователя
	// И добавить соответствующую запись в таблицу историй транзакций
	// Объединим эти операции в транзакцию
	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	// Обновление счетов
	endDeposit := startDeposit - amount
	endReserve := startReserve + amount
	_, err = tx.Exec(
		"UPDATE users SET deposit_account=$1, reserve_account=$2 WHERE user_id=$3",
		endDeposit,
		endReserve,
		userID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Отрицательное значение передается при снятии денег
	// Так как в таблице есть ограничение на то, что сумма только положительная
	// Умножим ее на минус 1
	if amount < 0 {
		amount *= -1
	}
	_, err = tx.Exec(
		"INSERT INTO transaction (user_id, datetime, amount, start_deposit, end_deposit, event, message)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userID, time.Now(), amount, startDeposit, endDeposit, "RESERVE", msg,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// ConfessionOrder подтверждает снятие amount денег пользователя с userID за услугу с serveceID и заказ с serviceID
func (b *AccountPostgres) ConfessionOrder(userID int64, serviceID int64, orderID int64, amount float64, details string) error {

	account := &model.User{}
	err := b.db.Get(account, "SELECT * FROM users WHERE user_id=$1", userID)

	// Проверим, достаточно ли у пользователя денег на резервном счете
	if amount > account.ReserveAccount {
		return errors.New("not enough money in the bank account")
	}

	// Cформируем сообщение для записи в таблицу транзакций
	msg := fmt.Sprintf("take %.2f on account", amount)
	if len(details) != 0 {
		msg += ": " + details
	}

	// Кейс : у пользователя уже есть запись в таблице
	// значит надо обновить это значение
	if err != nil {
		return b.confession(userID, amount, msg, serviceID, orderID, account.DepositAccount, account.ReserveAccount)
	}
	return err
}

// confession подтверждает снятие amount денег пользователя с userID за услугу с serveceID и заказ с serviceID
func (b *AccountPostgres) confession(userID int64, amount float64, msg string, serviceID int64, orderID int64, startDeposit float64, startReserve float64) error {
	// Надо сделать 2 вещи: снять с резервного счета деньги
	// И добавить соответствующую запись в таблицу историй транзакций
	// Объединим эти операции в транзакцию
	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	// Обновление счетов
	endReserve := startReserve - amount
	_, err = tx.Exec(
		"UPDATE users SET reserve_account=$1 WHERE user_id=$2",
		endReserve,
		userID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(
		"INSERT INTO transaction (user_id, datetime, amount, start_deposit, end_deposit, event, message)"+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		userID, time.Now(), amount, startDeposit, startDeposit, "TAKE", msg,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// TransferMoneyBetweenUsers переводит amount денег со счета пользователя с fromUserID пользователю с toUserID
func (b *AccountPostgres) TransferMoneyBetweenUsers(fromUserID int64, toUserID int64, amount float64, details string) error {

	// Достаем из БД запись о банковском счете пользователя-отправителя
	fromBankAccount := &model.User{}
	err := b.db.Get(fromBankAccount, `SELECT * FROM users WHERE user_id=$1`, fromUserID)

	// У пользователя еще нет записи баланса в таблице
	// Ошибка, так как снимать не с чего
	if err != nil {
		return errors.New("sender user does not have a bank account")
	}

	// Проверяем, достаточно ли денег у отправителя на счету
	if fromBankAccount.DepositAccount < amount {
		return errors.New("the sender does not have enough money in the account")
	}

	// Сформируем сообщение для записи в таблицу транзакций для отправителя
	fromMsg := fmt.Sprintf("transfer between accounts: withdraw %.2f from account", amount)
	if len(details) != 0 {
		fromMsg += ": " + details
	}

	// Сформируем сообщение для записи в таблицу транзакций для получателя
	toMsg := fmt.Sprintf("transfer between accounts: deposit %.2f to account", amount)
	if len(details) != 0 {
		toMsg += ": " + details
	}

	// Достаем из БД запись о банковском счете пользователя-получателя
	toBankAccount := &model.User{}
	err = b.db.Get(toBankAccount, `SELECT * FROM users WHERE user_id=$1`, toUserID)

	// Кейс
	// У пользователей уже есть запись баланса в таблице
	// Значит надо снять у отправителя, а получателю увеличить баланс
	return b.notFirstTransfer(fromUserID, toUserID, fromMsg, toMsg, fromBankAccount.DepositAccount, toBankAccount.DepositAccount, amount)
}

// notFirstTransfer реализует не первый перевод между пользователями
func (b *AccountPostgres) notFirstTransfer(
	fromUserID int64, toUserID int64,
	fromMsg string, toMsg string,
	fromStartBalance float64, toStartBalance float64,
	amount float64) error {

	// Надо сделать 4 вещи:
	// 1. Снять со счета отправителя
	// 2. Добавить соответствующую запись в таблицу историй транзакций
	// 3. Положить на счет получателю
	// 4. Добавить соответствующую запись в таблицу историй транзакций
	// Объединим эти операции в транзакцию

	tx, err := b.db.Beginx()
	if err != nil {
		return err
	}

	// 1. Снять со счета отправителя
	endBalance := fromStartBalance - amount
	_, err = tx.Exec(
		`UPDATE users SET deposit_account=$1 WHERE user_id=$2`,
		endBalance,
		fromUserID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 2. Добавить соответствующую запись в таблицу историй транзакций
	_, err = tx.Exec(
		`INSERT INTO transaction (user_id, datetime, amount, start_deposit, end_deposit, event, message) `+
			`VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		fromUserID, time.Now(), amount, fromStartBalance, endBalance, "TRANSFER", fromMsg,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 3. Положить на счет получателю
	endBalance = toStartBalance + amount
	_, err = tx.Exec(
		`UPDATE users SET deposit_account=$1 WHERE user_id=$2`,
		endBalance,
		toUserID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 4. Добавить соответствующую запись в таблицу историй транзакций
	_, err = tx.Exec(
		`INSERT INTO transactions_history (user_id, datetime, amount, start_deposit, end_deposit, event, message) `+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		toUserID, time.Now(), amount, toStartBalance, endBalance, "TRANSFER", toMsg,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetTransactionsHistory возвращает историю транзакций пользователя с userID
func (b *AccountPostgres) GetTransactionsHistory(userID int64) ([]*model.Transactions, error) {
	history := make([]*model.Transactions, 0)
	err := b.db.Select(&history, `SELECT * FROM transaction WHERE user_id=$1`, userID)
	if err != nil {
		return nil, err
	}
	return history, nil
}
