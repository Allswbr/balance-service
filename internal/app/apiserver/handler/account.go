package handler

// Context содержит детали запроса, проверяет и сериализует JSON ...
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Просмотр баланса
func (h *Handler) getUserAccountBalance(c *gin.Context) {
	// Тело запроса
	type Request struct {
		UserID int64 `json:"user_id"`
	}

	// Декодирование тела запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверяем существование пользователя в системе
	_, err := h.Services.GetUserByID(req.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user doesn't exist in the system: %s", err.Error()))
		return
	}

	var userBalance map[string]string

	res, err := h.Services.GetBalanceByUserID(req.UserID)
	if err != nil {
		// Проверяем существование пользователя в системе
		if len(res) == 0 {
			newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user does not have a bank account: : %s", err.Error()))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with getting user balance: %s", err.Error()))
		}
		return
	}
	userBalance = res

	c.JSON(http.StatusOK, userBalance)
}

// Пополнение
func (h *Handler) depositUserAccount(c *gin.Context) {

	// Тело запроса
	type Request struct {
		UserID  int64   `json:"user_id"`
		Amount  float64 `json:"amount"`
		Details string  `json:"details"`
	}

	// Декодирование тела запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверка существования пользователя в системе
	_, err := h.Services.GetUserByID(req.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user does not exist in the system: %s", err.Error()))
		return
	}

	err = h.Services.DepositMoneyToUser(req.UserID, req.Amount, req.Details)

	if err != nil {
		if strings.Contains(err.Error(), "the amount must be greater than zero") {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error with update user balance: %s", err.Error()))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with update user balance: %s", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("%.2f were successfully deposited to user bank account", req.Amount),
	})
}

// Резервирование
func (h *Handler) reservationUserAccount(c *gin.Context) {

	// Тело запроса
	type Request struct {
		UserID    int64   `json:"user_id"`
		ServiceID int64   `json:"service_id"`
		OrderID   int64   `json:"order_id"`
		Amount    float64 `json:"amount"`
		Details   string  `json:"details"`
	}

	// Декодирование тела запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверка существования пользователя в системе
	_, err := h.Services.GetUserByID(req.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user does not exist in the system: %s", err.Error()))
		return
	}

	err = h.Services.ReservationUserAccount(req.UserID, req.ServiceID, req.OrderID, req.Amount, req.Details)

	if err != nil {
		if strings.Contains(err.Error(), "the amount must be greater than zero") {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error with update user balance: %s", err.Error()))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with update user balance: %s", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("%.2f were successfully reserve on user bank account", req.Amount),
	})
}

// Подтверждение выручки
func (h *Handler) confessionOrder(c *gin.Context) {

	// Тело запроса
	type Request struct {
		UserID    int64   `json:"user_id"`
		ServiceID int64   `json:"service_id"`
		OrderID   int64   `json:"order_id"`
		Amount    float64 `json:"amount"`
		Details   string  `json:"details"`
	}

	// Декодирование тела запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверка существования пользователя в системе
	_, err := h.Services.GetUserByID(req.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user does not exist in the system: $s", err.Error()))
		return
	}

	err = h.Services.ConfessionOrder(req.UserID, req.ServiceID, req.OrderID, req.Amount, req.Details)

	if err != nil {
		if strings.Contains(err.Error(), "the amount must be greater than zero") {
			newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error with update user balance: %s", err.Error()))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with update user balance: %s", err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("%.2f were successfully take from user bank account", req.Amount),
	})
}

// Перевод
func (h *Handler) transferMoneyBetweenUsers(c *gin.Context) {
	type Request struct {
		FromUserID int64   `json:"from_user_id"`
		ToUserID   int64   `json:"to_user_id"`
		Amount     float64 `json:"amount"`
		Details    string  `json:"details"`
	}

	// Декодируем тело запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверяем существование юзера-отправителя
	_, err := h.Services.GetUserByID(req.FromUserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("sender user does not exist in the system: %s", err.Error()))
		return
	}

	// Проверяем существования юзера-получателя
	_, err = h.Services.GetUserByID(req.ToUserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("destination user does not exist in the system %s", err.Error()))
		return
	}

	// Делегируем работу слою сервиса
	err = h.Services.TransferMoneyBetweenUsers(req.FromUserID, req.ToUserID, req.Amount, req.Details)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with transfer money between users: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("%.2f were successfully transferred between users", req.Amount),
	})

}

// История транзакций
func (h *Handler) getUserTransactionHistory(c *gin.Context) {
	type Request struct {
		UserID int64 `json:"user_id"`
	}

	// Декодируем тело запроса
	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	// Проверяем существования пользователя в системе
	_, err := h.Services.GetUserByID(req.UserID)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("this user does not exist in the system: %s", err.Error()))
		return
	}

	// Получаем историю транзакций
	history, err := h.Services.GetTransactionsHistory(req.UserID)
	if len(history) == 0 {
		newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("there is no transaction history for this user: %s", err.Error()))
		return
	}

	// Сортируем результат
	sortParam := c.Query("sort")
	switch strings.ToLower(sortParam) {
	case "amount":
		sort.Slice(history, func(i, j int) bool {
			return history[i].Amount < history[j].Amount
		})
	case "date", "":
		sort.Slice(history, func(i, j int) bool {
			return history[i].Datetime.Unix() < history[j].Datetime.Unix()
		})
	default:
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("invalid sort param in query: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, history)
}

// Отчёт бухгалтерии за месяц-год
func (h *Handler) getOrdersYearMonth(c *gin.Context) {
	type Request struct {
		Year  int64 `json:"year"`
		Month int64 `json:"month"`
	}
}

// Достаем id пользователя из url
func (h *Handler) getUserID(c *gin.Context) (int64, error) {
	strUserID := c.Param("user_id")
	intUserID, err := strconv.Atoi(strUserID)
	if err != nil {
		return 0, err
	}
	return int64(intUserID), nil
}
