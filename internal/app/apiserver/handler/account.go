package handler

// Context содержит детали запроса, проверяет и сериализует JSON ...
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Просмотр баланса
func (h *Handler) getUserAccountBalance(c *gin.Context) {
	type Request struct {
		UserID int64 `json:"user_id"`
	}

	req := &Request{}
	if err := c.BindJSON(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("incorrect request body: %s", err.Error()))
		return
	}

	var userBalance map[string]string

	res, err := h.Services.GetBalanceByUserID(req.UserID)
	if err != nil {
		if len(res) == 0 {
			newErrorResponse(c, http.StatusNotFound, "this user does not have a bank account")
		} else {
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error with getting user balance: %s", err.Error()))
		}
		return
	}
	userBalance = res

	c.JSON(http.StatusOK, map[string]interface{}{
		"deposit": userBalance["deposit"],
		"reserve": userBalance["reserve"],
	})
}

// Пополнение
func (h *Handler) depositUserAccount(c *gin.Context) {
	type Request struct {
		UserID int64  `json:"user_id"`
		Amount string `json:"amount"`
	}
}

// Резервирование
func (h *Handler) reservationUserAccount(c *gin.Context) {
	type Request struct {
		UserID    int64  `json:"user_id"`
		ServiceID int64  `json:"service_id"`
		OrderID   int64  `json:"order_id"`
		Amount    string `json:"amount"`
	}
}

// Подтверждение выручки
func (h *Handler) confessionOrder(c *gin.Context) {
	type Request struct {
		UserID    int64  `json:"user_id"`
		ServiceID int64  `json:"service_id"`
		OrderID   int64  `json:"order_id"`
		Amount    string `json:"amount"`
	}
}

// Перевод
func (h *Handler) transferMoneyBetweenUsers(c *gin.Context) {
	type Request struct {
		FromUserID int64  `json:"from_user_id"`
		ToUserID   int64  `json:"to_user_id"`
		Amount     string `json:"amount"`
	}
}

// История транзакций
func (h *Handler) getUserTransactionHistory(c *gin.Context) {
	type Request struct {
		UserID int64 `json:"user_id"`
	}
}

// Отчёт бухгалтерии за месяц-год
func (h *Handler) getOrdersYearMonth(c *gin.Context) {
	type Request struct {
		Year  int64 `json:"year"`
		Month int64 `json:"month"`
	}
}
