package Controllers

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRole(c *gin.Context) {
	var input struct {
		RoleUID  string `json:"role_uid" binding:"required"`
		NameRole string `json:"name_role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}

	role := Models.TURole{
		RoleUID:  input.RoleUID,
		NameRole: input.NameRole,
	}

	if err := Database.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Status":  "Success",
		"Message": "Role Created Successfully",
		"Data":    role,
	})
}

func GetAllRoles(c *gin.Context) {
	var roles []Models.TURole
	if err := Database.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "All Roles",
		"Data":    roles,
	})
}

func GetRoleById(c *gin.Context) {
	var role Models.TURole
	if err := Database.DB.First(&role, c.Param("role_uid")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Role",
		"Data":    role,
	})
}

func UpdateRole(c *gin.Context) {
	var role Models.TURole
	if err := Database.DB.First(&role, c.Param("role_uid")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}
	if err := Database.DB.Save(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Role Updated Successfully",
		"Data":    role,
	})
}

func DeleteRole(c *gin.Context) {
	var role Models.TURole
	if err := Database.DB.Where("role_uid = ?", c.Param("role_uid")).First(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}

	if err := Database.DB.Delete(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"Message": "Akses Ditolak Server",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Role Deleted Successfully",
		"Data":    role,
	})
}
