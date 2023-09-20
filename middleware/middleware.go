package middleware

import (
	"log"
	"net/http"
	"qkeruen/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	"fmt"
)

func AuthorizeJWTDriver(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No token found."})
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["role"] != "driver" {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "wrong type role"})
				return
			}
			log.Println(claims["phone_number"])
			log.Println(claims["role"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "unauthorized"})
		}

	}
}

func AuthorizeJWTUser(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("here")
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No toke found."})
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["role"] != "user" {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "wrong type role"})
				return
			}
			log.Println(claims["phone_number"])
			log.Println(claims["role"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "unauthorized"})
		}

	}
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Content-Length", "Authorization", "X-CSRF-Token", "Content-Type", "Accept", "X-Requested-With", "Bearer"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type", "application/json", "Content-Length", "Accept-Encoding", "X-CSRF-Token",  "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://api.qkeruen.kz"
		},
		MaxAge: 12 * time.Hour,
	})
}
