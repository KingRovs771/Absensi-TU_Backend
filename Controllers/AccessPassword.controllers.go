package Controllers

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAccessPasswords(c *gin.Context){
	var input struct{
		PasswordAccess string `json:password_access`
	}

	if err := c.ShouldBindBodyWithJSON(&input); err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"Status" : "Error",
			"Message" : "Bad Request",
			"Error" : err.Error(),
		})
		return
	}

	accessPassword := Models.TUAccessPasswords{
		PasswordAccess : input.PasswordAccess,
	}

	if err := Database.DB.Create(&accessPassword).Error; err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Internal Server Error",
			"Error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Membuat Password Access",
		"Data" : accessPassword,
	})
}

func GetAllAccessPasswords(c *gin.Context){
	var accesspassword []Models.TUAccessPasswords

	if err := Database.DB.Find(&accesspassword).Error; err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Internal Server Error",
			"Error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Mendapatkan Data Access Password",
		"Data" : accesspassword,
	})
}

func GetUIDAccessPasswords(c *gin.Context){
	UIDAccessPassword := c.Param("uid")
	var accesspassword Models.TUAccessPasswords

	if err := Database.DB.Where("access_passwords_uid = ?", UIDAccessPassword).First(accesspassword).Error; err !=nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{
				"Status" : "Error",
				"Message" : "Data Not Found",
				"Error" : err.Error(),
			})
			return
		} 
		c.JSON(http.StatusBadRequest, gin.H{
			"Status" : "Error",
			"Message" : "Bad Request",
			"Error" : err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Mendapatkan Data UID Access Password",
		"Data" : accesspassword,
	})
}

func UpdateAccessPasswords(c *gin.Context){

}

func DeleteAccessPasswords(c *gin.Context){
	UIDAccessPassword := c.Param("uid")

	deleteAccess := Database.DB.Where("access_passwords_uid = ?", UIDAccessPassword).Delete(&Models.TUAccessPasswords{})

	if deleteAccess.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Internal Server Error",
		})
		return
	}

	if deleteAccess.RowsAffected == 0{
		c.JSON(http.StatusNotFound, gin.H{
			"Status" : "Error",
			"Message" : "Not Found",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Menghapus Data Access Password",
		"Data" : deleteAccess,
	})
}