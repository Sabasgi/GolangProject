package models

import "github.com/golang-jwt/jwt"

type JWTBasicConfig struct {
	Username string `json:"username"`
	UserId   int    `json:"user_id"`
	Role     string `json:"role"`
	LabID    int    `json:"lab_id"`
	jwt.StandardClaims
}
