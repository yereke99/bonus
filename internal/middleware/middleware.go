package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://api.qkeruen.kz"
		},
		MaxAge: 12 * time.Hour,
	})
}
