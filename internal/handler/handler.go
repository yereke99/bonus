package handler

import (
	"bonus/config"
	"bonus/internal/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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

	/*
		r.Use(cors.New(cors.Config{
				AllowOrigins:     []string{"*"},
				AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer", "Authority"},
				ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
				AllowCredentials: true,
				AllowOriginFunc:  func(origin string) bool { return origin == "https://api.qkeruen.kz" },
			}))
	*/

	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	r.POST("/api/v1/code", h.SendCode)
	r.POST("/api/v1/registry", h.Registry)
	r.POST("/api/v1/login", h.Login)
	r.POST("/api/v1/refresh", h.RefreshToken)

	r.POST("/api/v1/company", h.CreateCompany)

	return r
}

// AuthorizeJWTUser is a middleware function for authorizing JWT users
func (h *Handler) AuthorizeJWTUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			h.zapLogger.Error("No token found in header")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No token found."})
			return
		}

		// Extract the Bearer token from the Authorization header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			h.zapLogger.Error("Invalid Authorization header format")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Authorization header format."})
			return
		}

		tokenString := tokenParts[1]

		token, err := h.service.JWTService.ValidateToken(tokenString)
		if err != nil {
			h.zapLogger.Error("Error validating token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role, ok := claims["role"].(string); ok {
				if role != "user" {
					h.zapLogger.Error("Wrong type role", zap.String("role", role))
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Wrong type role"})
					return
				}
				h.zapLogger.Info("user role", zap.Any("role", role))
				c.Next()
			} else {
				h.zapLogger.Error("Invalid token", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
		}
	}
}

// AuthorizeJWTAdmin is a middleware function for authorizing JWT admins
func (h *Handler) AuthorizeJWTAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			h.zapLogger.Error("No token found in header")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No token found."})
			return
		}

		// Extract the Bearer token from the Authorization header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			h.zapLogger.Error("Invalid Authorization header format")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Authorization header format."})
			return
		}

		tokenString := tokenParts[1]

		token, err := h.service.JWTService.ValidateToken(tokenString)
		if err != nil {
			h.zapLogger.Error("Error validating token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role, ok := claims["role"].(string); ok {
				if role != "admin" {
					h.zapLogger.Error("Wrong type role", zap.String("role", role))
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Wrong type role"})
					return
				}
				h.zapLogger.Info("admin role", zap.Any("role", role))
				c.Next()
			} else {
				h.zapLogger.Error("Invalid token", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
		}
	}
}

// AuthorizeJWTPartner is a middleware function for authorizing JWT partners
func (h *Handler) AuthorizeJWTPartner() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			h.zapLogger.Error("No token found in header")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No token found."})
			return
		}

		// Extract the Bearer token from the Authorization header
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			h.zapLogger.Error("Invalid Authorization header format")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid Authorization header format."})
			return
		}

		tokenString := tokenParts[1]

		token, err := h.service.JWTService.ValidateToken(tokenString)
		if err != nil {
			h.zapLogger.Error("Error validating token", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if role, ok := claims["role"].(string); ok {
				if role != "partner" {
					h.zapLogger.Error("Wrong type role", zap.String("role", role))
					c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Wrong type role"})
					return
				}
				h.zapLogger.Info("partner role", zap.Any("role", role))
				c.Next()
			} else {
				h.zapLogger.Error("Invalid token", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			}
		}
	}
}
