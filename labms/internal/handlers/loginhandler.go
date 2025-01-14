package handlers

import (
	"log"
	"net/http"
	"os"
	models "repogin/internal/models"
	masters "repogin/internal/models/masters"
	"repogin/logs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func (h *MainHandlers) LoginRoute(c *gin.Context) {

	var user masters.Userr
	bindError := c.Bind(&user)
	if bindError != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"error": bindError.Error()})
		logs.Error("ERROR : LoginRoute", bindError)
	}
	userDetails, err := h.Uservice.GetUserLoginService(user)
	if err != nil {
		logs.Error("ERROR : LoginRoute", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	t, terr := GenerateToken(userDetails)
	if terr != nil {
		logs.Error("ERROR : LoginRoute", terr)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": terr.Error()})
		return
	}
	logs.Debug("Login Route Toke : ", t)
	c.Request.Header.Set("Authorization", t)
	c.JSON(http.StatusOK, gin.H{"token": t, "role": userDetails.Role, "message": "SuccessFull login"})
	return
}

func GenerateToken(u masters.Userr) (string, error) {
	claims := models.JWTBasicConfig{
		Username: u.Username,
		UserId:   u.UserID,
		Role:     u.Role,
		LabID:    u.LabID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := os.Getenv("key")
	log.Println("kkeyyy ", key)
	t, terr := token.SignedString([]byte(key))
	if terr != nil {
		log.Println("ERROR : GENERATE TOKEN ", terr)
		return "", terr
	}
	return t, nil
}
