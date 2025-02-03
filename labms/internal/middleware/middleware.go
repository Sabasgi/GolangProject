package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"repogin/internal/models"
	masters "repogin/internal/models/masters"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func TokenValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		arr := strings.Split(c.GetHeader("Authorization"), " ")
		// fmt.Println("arrrr,,header", arr, c.GetHeader("Authorization"))
		if len(arr) > 1 {
			token := arr[1]
			if token == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
				c.Abort()
				return
			}
			user, err := validateToken(token)
			if err != nil {
				log.Println("errror in token validation ", err)
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				c.Abort()
				return
			}
			// Store user details in context
			c.Set("role", user.Role)
			c.Set("userId", user.UserID)
			c.Set("labId", user.LabID)
			c.Set("username", user.Username)
			c.Next()

		} else {
			fmt.Println("c.Hedaer", arr)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return

		}
		// Validate token (pseudo-code; implement as needed)

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

func RoleAuthorization(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.GetString("role") // Assume role is extracted from token in earlier middleware

		// Super admin bypasses all checks
		if userRole == "superadmin" {
			c.Next()
			return
		}

		// Check if the role is allowed for this API
		if !contains(allowedRoles, userRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Helper function to check if a role is in the allowed list
func contains(slice []string, item string) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
func validateToken(tokenString string) (masters.Userr, error) {
	// Parse the JWT token
	claims := &models.JWTBasicConfig{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Return the signing key
		return []byte(os.Getenv("key")), nil
	})
	if err != nil {
		return masters.Userr{}, errors.New("invalid token")
	}

	// Check if the token is valid
	if !token.Valid {
		return masters.Userr{}, errors.New("invalid or expired token")
	}

	// Check token expiration
	if claims.ExpiresAt < time.Now().Unix() {
		return masters.Userr{}, errors.New("token expired")
	}

	// Map the claims to the Userr struct
	user := masters.Userr{
		UserID: claims.UserId,
		LabID:  claims.LabID,
		// Name:      claims.Name,
		// Email:     claims.Email,
		Role:      claims.Role,
		Username:  claims.Username,
		CreatedAt: time.Unix(claims.IssuedAt, 0).String(),
	}

	return user, nil
}
