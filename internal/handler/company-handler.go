package handler

import (
	"bonus/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCompany(c *gin.Context) {
	var company *domain.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	resp, err := h.service.CompanyService.CreateCompany(company)
	if err != nil {
		c.JSON(
			http.StatusConflict, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
