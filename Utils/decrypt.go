package Utils

import (
	"absensiTU-Backend/Models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"strings"
	"time"
)

type CustomeClaimsJWT struct {
	UsersUID string `json:"users_uid"`
	RoleUID  string `json:"role_uid"`
	FullName string `json:"full_name"`
	jwt.RegisteredClaims
}

func GenerateJWTAdminTU(Users Models.TUUsers) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	tokenLifespanStr := os.Getenv("JWT_LIFESPAN")
	tokenLifespan, err := strconv.Atoi(tokenLifespanStr)

	if err != nil {
		tokenLifespan = 1
	}

	claims := CustomeClaimsJWT{
		UsersUID: Users.UserUID,
		RoleUID:  Users.RoleUID,
		FullName: Users.FullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tokenLifespan))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}

func ValidateJWT(c *gin.Context) (*CustomeClaimsJWT, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("JWT_SECRET belum di-set")
	}

	tokenString := getTokenFromRequest(c)
	if tokenString == "" {
		return nil, errors.New("Token otorisasi tidak disediakan")
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomeClaimsJWT{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Metode signing tidak terduga")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomeClaimsJWT); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Token tidak Valid Silakan Ulangi Lagi")
}
