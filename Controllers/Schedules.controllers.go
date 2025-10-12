package Controllers

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SchedulesInput struct{
	Day 		string 	`json:"day" binding:"required"`
	StartTime 	string	`json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime 	string 	`json:"end_time" binding:"required,datetime=15:04:05"`
}

type SchdulesUpdate struct{
	Day 		string 	`json:"day" binding:"required"`
	StartTime 	string	`json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime 	string 	`json:"end_time" binding:"required,datetime=15:04:05"`
}

func GetAllSchedules(c *gin.Context){
	var schedules []Models.TUSchedules

	if err := Database.DB.Find(&schedules).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Ditolak Server",
			"Error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message" : "Sukses Mendapatkan Data Server",
		"Data" : schedules,
	})
}

func CreateSchdules(c *gin.Context){
	var input SchedulesInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Invalid Input",
			"Error" : err.Error(),
		})
		return
	}

	schdules := Models.TUSchedules{
		ScheduleUID: uuid.New().String(),
		Day:         input.Day,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
	}

	if err := Database.DB.Create(&schdules).Error;err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message" : "Ditolak Server",
			"Error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Membuat Jadwal Baru",
		"Data" : schdules,
	})
}

func GetByUIDSchdules(c *gin.Context){
	uid := c.Param("uid")
	var schedule Models.TUSchedules

	if err := Database.DB.Where("schedule_uid = ?", uid).First(&schedule).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Status" : "Not Found",
				"Message": "Schedule not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message": "Database error", 
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message": "Schedule retrieved successfully",
		"Data":    schedule,
	})
}

func UpdateSchdules(c *gin.Context){
	uid := c.Param("uid")
	var existingSchedule Models.TUSchedules

	if err := Database.DB.Where("schedule_uid = ?", uid).First(&existingSchedule).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Status" : "Error",
				"Message": "Schedule not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message": "Database error", 
			"Error": err.Error(),
		})
		return
	}

	var input SchdulesUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Status" : "Error",
			"Message": "Invalid input", 
			"Error": err.Error(),
		})
		return
	}

	existingSchedule.Day = input.Day
	existingSchedule.StartTime = input.StartTime
	existingSchedule.EndTime = input.EndTime

	if err := Database.DB.Save(&existingSchedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message": "Failed to update schedule", 
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message": "Schedule updated successfully",
		"Data":    existingSchedule,
	})
}

func DeleteSchdules(c *gin.Context){
	uid := c.Param("uid")

	result := Database.DB.Where("schedule_uid = ?", uid).Delete(&Models.TUSchedules{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message": "Failed to delete schedule", 
			"Error": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Status" : "Error",
			"Message": "Schedule not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Status" : "Success",
		"Message" : "Berhasil Menghapus Data Schedules",
		"Data" : result,
	})
}