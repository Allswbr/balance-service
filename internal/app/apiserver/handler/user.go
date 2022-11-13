package handler

import (
	"fmt"
	"github.com/Allswbr/balance-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Регистрация нового пользователя
func (h *Handler) createUser(c *gin.Context) {
	user := &model.User{}

	id, err := h.Services.CreateUser(user)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error creating user: %s", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

// Просмотр всех существующих пользователей
func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.Services.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error getting users: %s", err.Error()))
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "no users yet",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
