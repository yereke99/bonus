package handler

import (
	"bonus/internal/domain"
	"bonus/traits"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCompany(c *gin.Context) {
	var company *domain.CompanyRequest
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	resp, err := h.service.CompanyService.CreateCompany(domain.CompanyRequest)
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

// GetCompanies godoc
// @Summary Get list of companies
// @Description Retrieves a list of all companies.
// @Tags company
// @Produce json
// @Success 200 {array} domain.Company "List of companies"
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company [get]
func (h *Handler) GetCompanies(c *gin.Context) {
	companies, err := h.service.CompanyService.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// NotifyUser godoc
// @Summary Send a notification to the user
// @Description Sends a notification related to the user’s account activities.
// @Tags user
// @Accept json
// @Produce json
// @Param notification body domain.NotificationRequest true "Notification Data"
// @Success 200 {object} map[string]string{"status": "Notification sent"}
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/user/notify [post]
func (h *Handler) NotifyUser(c *gin.Context) {
	var req domain.NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Dummy implementation for sending notification
	// Implement actual notification logic here
	/*
		if err := h.service.UserService.SendNotification(req); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send notification"})
				return
			}
	*/

	c.JSON(http.StatusOK, gin.H{"status": "Notification sent"})
}

// CalculateBonus godoc
// @Summary Calculate bonus
// @Description Calculates the bonus based on the percentage of the product or service cost.
// @Tags company
// @Accept json
// @Produce json
// @Param bonus body domain.BonusCalculationRequest true "Bonus Calculation Data"
// @Success 200 {object} domain.BonusCalculationResponse "Successfully calculated bonus"
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/bonus [post]
func (h *Handler) CalculateBonus(c *gin.Context) {
	var req domain.BonusCalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		resp, err := h.service.CompanyService.CalculateBonus(&req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate bonus"})
				return
			}
	*/

	c.JSON(http.StatusOK, "ok")
}

// AddBarcode godoc
// @Summary Add barcode
// @Description Links the barcode identifier to the bonus size in the Partner's database.
// @Tags company
// @Accept json
// @Produce json
// @Param barcode body domain.BarcodeRequest true "Barcode Data"
// @Success 200 {object} map[string]string{"status": "Barcode linked successfully"}
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/add-code [post]
func (h *Handler) AddBarcode(c *gin.Context) {
	var req domain.BarcodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		if err := h.service.CompanyService.AddBarcode(&req); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to link barcode"})
				return
			}
	*/

	c.JSON(http.StatusOK, gin.H{"status": "Barcode linked successfully"})
}

// CalculateCommission godoc
// @Summary Calculate commission
// @Description Calculates the commission amount and displays it in whole numbers in monitoring WBS and Partner sales.
// @Tags company
// @Accept json
// @Produce json
// @Param commission body domain.CommissionCalculationRequest true "Commission Calculation Data"
// @Success 200 {object} domain.CommissionCalculationResponse "Successfully calculated commission"
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/calculate-commission [post]
func (h *Handler) CalculateCommission(c *gin.Context) {
	var req domain.CommissionCalculationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		resp, err := h.service.CompanyService.CalculateCommission(&req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate commission"})
				return
			}
	*/

	c.JSON(http.StatusOK, nil)
}

// DoubleBonus godoc
// @Summary Double bonus balance
// @Description Doubles the bonus balance of the Partner.
// @Tags company
// @Accept json
// @Produce json
// @Param doubleBonus body domain.DoubleBonusRequest true "Double Bonus Data"
// @Success 200 {object} map[string]string{"status": "Bonus doubled successfully"}
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/increase-bonus [post]
func (h *Handler) DoubleBonus(c *gin.Context) {
	var req domain.DoubleBonusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*
		if err := h.service.CompanyService.DoubleBonus(&req); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to double bonus"})
				return
			}
	*/

	c.JSON(http.StatusOK, gin.H{"status": "Bonus doubled successfully"})
}

// SearchCompanies godoc
// @Summary Search for companies
// @Description Searches for companies based on query parameters like company name, IIN, or city.
// @Tags company
// @Produce json
// @Param name query string false "Company name"
// @Param iin query string false "Company IIN"
// @Param city query string false "City"
// @Success 200 {array} domain.Company "List of matching companies"
// @Failure 400 {object} map[string]string{"error": "Bad request"}
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/searching [get]
func (h *Handler) SearchCompanies(c *gin.Context) {
	name := c.Query("name")
	iin := c.Query("iin")
	city := c.Query("city")

	companies, err := traits.SearchCompanies(name, iin, city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// MonitorBonus godoc
// @Summary Monitor bonuses and sales
// @Description Displays the total amounts of bonuses issued/received, sales, and commission amounts.
// @Tags company
// @Produce json
// @Success 200 {object} domain.BonusMonitoringResponse "Monitoring data for bonuses and sales"
// @Failure 500 {object} map[string]string{"error": "Internal Server Error"}
// @Router /api/v1/company/monitoring-bonus [get]
func (h *Handler) MonitorBonus(c *gin.Context) {
	monitoringData, err := traits.MonitorBonus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch monitoring data"})
		return
	}

	c.JSON(http.StatusOK, monitoringData)
}
