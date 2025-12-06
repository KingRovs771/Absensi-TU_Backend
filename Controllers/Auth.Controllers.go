package Controllers

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"absensiTU-Backend/Utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status":  " Error",
			"Message": "Input Tidak Valid",
		})
		return
	}

	var user Models.TUUsers
	if err := Database.DB.Where("email = ?", strings.ToLower(input.Email)).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Status":  " Error",
				"Message": "Email dan Password Salah",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status":  " Error",
			"Message": "Gagal Memproses Permintaan",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}

	token, err := Utils.GenerateJWTAdminTU(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "Login berhasil",
		"token":     token,
		"user_uid":  user.UsersUID,
		"full_name": user.FullName,
		"role_uid":  user.RoleUID,
	})
}

func LogoutHandler(c *gin.Context) {
	userUID, exits := c.Get("users_uid")

	if !exits {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Status":  "Error",
			"Message": "Token Tidak Valid",
		})
		return
	}

	var users Models.TUUsers
	if err := Database.DB.Where("user_uid = ?", userUID.(string)).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  "Error",
			"Message": "Users Not Found",
			"Error":   err.Error(),
		})
	}
}
