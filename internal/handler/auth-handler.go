package handler

import (
	"bonus/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendCode(c *gin.Context) {
	var sign domain.Registry

	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	if err := h.service.AuthService.SendCode(&sign); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	err := h.service.AuthService.SendCode(&sign)
	if err != nil {
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

	resp, err := h.service.AuthService.Registry(&registry)
	if err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

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
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err,
			},
		)
		return
	}

	c.JSON(http.StatusOK, resp)
}
