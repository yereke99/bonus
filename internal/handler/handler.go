package handler

import (
	"bonus/config"
	"bonus/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service   *service.Services
	zapLogger *zap.Logger
	appConfig *config.Config
}

func NewHandler(service *service.Services, zapLogger *zap.Logger, appConfig *config.Config) *Handler {
	return &Handler{
		service:   service,
		zapLogger: zapLogger,
		appConfig: appConfig,
	}
}

func (h *Handler) InitHandler() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/api/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	return r
}
