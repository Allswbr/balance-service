package handler

import (
	"github.com/Allswbr/balance-service/internal/app/service"
	"github.com/gin-gonic/gin"
)

// Handler - структура для обработчика
type Handler struct {
	Services *service.Service
}

// NewHandler - конструктор для обработчика
func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

// InitRoutes - функция для инициализации обработчиков эндпоинтов
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	// Основные роуты для работы с балансом
	balance := router.Group("/balance")
	{
		balance.GET("", h.getUserAccountBalance)                  // Просмотр баланса
		balance.POST("/deposit", h.depositUserAccount)            // Пополнение
		balance.POST("/reservation", h.reservationUserAccount)    // Резервирование
		balance.POST("/confession", h.confessionOrder)            // Подтверждение выручки
		balance.POST("/transfer", h.transferMoneyBetweenUsers)    // Перевод
		balance.GET("/transactions", h.getUserTransactionHistory) // История транзакций
		balance.GET("/report-month-year", h.getOrdersYearMonth)   // Отчёт
	}

	users := router.Group("/users")
	{
		// Дополнительные роуты для удобства работы с пользователями
		users.POST("/create-user", h.createUser) // Регистрация нового пользователя
		users.GET("", h.getAllUsers)             // Просмотр всех существующих пользователей
	}

	return router
}
