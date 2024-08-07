package handler

import (
	"bonus/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCode(c *gin.Context) {
	var code domain.CodeRequest

	if err := c.ShouldBindJSON(&code); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	if err := h.service.AuthService.SendCode(&code); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, "sent")
}

func (h *Handler) Registry(c *gin.Context) {
	var registry domain.RegistryRequest

	if err := c.ShouldBindJSON(&registry); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	tokens, err := h.service.AuthService.Registry(&registry)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, tokens)
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

	c.JSON(http.StatusOK, "")
}
