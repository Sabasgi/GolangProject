package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"repogin/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func BaicAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			arr := strings.Split(c.Request().Header.Get("Authorization"), " ")
			fmt.Println("arrrr,,header", arr, c.Request().Header)
			if len(arr) > 1 {
				Auth := arr[1]
				token, terr := jwt.Parse(Auth, func(t *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("key")), nil
				})
				if terr != nil {
					fmt.Println("ERROR:TokenPrasingERROR", terr)
					return echo.ErrUnauthorized

				}
				if token.Valid {
					fmt.Println("token is valid", token.Valid)
					return next(c)
				} else {
					fmt.Println("token is NOT valid ", token.Valid)
					return echo.ErrUnauthorized
				}
			}
			fmt.Println("c.Hedaer", arr)
			return echo.ErrUnauthorized
			// return nil
		}
	}
}

// Middleware to validate JWT and check user role
func RoleBasedMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := strings.Split(c.GetHeader("Authorization"), " ")[1]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}
		log.Println("Token string", tokenString)
		// Parse and validate the token
		claims := &models.JWTBasicConfig{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the role matches the required role
		if claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: Insufficient permissions"})
			c.Abort()
			return
		}

		// If valid, store the claims in the context for later use
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("userId", claims.UserId)
		c.Next()
	}
}
