package handler

// Context содержит детали запроса, проверяет и сериализует JSON ...
import "github.com/gin-gonic/gin"

// Просмотр баланса
func (h *Handler) getUserAccountBalance(c *gin.Context) {

}

// Пополнение
func (h *Handler) depositUserAccount(c *gin.Context) {

}

// Резервирование
func (h *Handler) reservationUserAccount(c *gin.Context) {

}

// Подтверждение выручки
func (h *Handler) confessionOrder(c *gin.Context) {

}

// Перевод
func (h *Handler) transferMoneyBetweenUsers(c *gin.Context) {

}

// История транзакций
func (h *Handler) getUserTransactionHistory(c *gin.Context) {

}

// Отчёт бухгалтерии за месяц-год
func (h *Handler) getOrdersYearMonth(c *gin.Context) {

}
