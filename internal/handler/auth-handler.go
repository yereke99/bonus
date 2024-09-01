package handler

import (
	"bonus/internal/domain"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCode(c *gin.Context) {
	var sign domain.Registry

	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.service.AuthService.SendCode(&sign); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "sent")
}

func (h *Handler) Registry(c *gin.Context) {
	var registry domain.RegistryRequest

	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	resp, err := h.service.AuthService.Registry(&registry)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) UpdateRegistry(c *gin.Context) {
	var registry domain.RegistryRequest

	// Получаем ID из параметров маршрута
	idParam := c.Param("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "invalid user ID",
			},
		)
		return
	}

	// Привязываем тело запроса к структуре
	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Вызываем сервис для обновления пользователя
	resp, err := h.service.AuthService.UpdateUser(userID, &registry)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Возвращаем обновленные данные пользователя
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c *gin.Context) {
	var registry domain.Registry

	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	resp, err := h.service.AuthService.Login(&registry)
	if err != nil {
		if err.Error() == "user does not exist" {
			c.JSON(http.StatusOK, "code is valid, but user does not exist")
			return
		}
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) RefreshToken(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Authorization header is missing",
			},
		)
		return
	}

	// Split the "Bearer" prefix from the token
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": "Invalid Authorization header format",
			},
		)
		return
	}

	// The actual token is the second part
	token := tokenParts[1]

	// Call the service to refresh the token
	resp, err := h.service.JWTService.RefreshToken(token)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// Return the refreshed token in the response
	c.JSON(http.StatusOK, resp)
}
