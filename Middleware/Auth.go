package Middleware

import (
	"absensiTU-Backend/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequireAuth(c *gin.Context) {
	claims, err := Utils.ValidateJWT(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		c.Abort()
		return
	}
	c.Set("users_uid", claims.UsersUID)
	c.Next()
}
