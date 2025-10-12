package Controllers

import (
	"absensiTU-Backend/Database"
	"absensiTU-Backend/Models"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AttendanceDetailResponse struct {
	AttendancesUID        string    `json:"attendances_uid"`
	UsersUID              string    `json:"users_uid"`
	FullName              string    `json:"fullname"`
	UserProfilePictureURL string    `json:"user_profile_picture_url"`
	AttendanceDate        string    `json:"attendance_date"`
	CheckIn               time.Time `json:"check_in"`
	StatusIn              string    `json:"status_in"`
	CheckInImageURL       string    `json:"check_in_image_url"`
	CheckOut              time.Time `json:"check_out"`
	StatusOut             string    `json:"status_out"`
	CheckOutImageURL      string    `json:"check_out_image_url"`
}

func GetAttendancesByDate(c *gin.Context) {
	dateStr := c.Query("date")
	var targetDate time.Time
	var err error

	if dateStr == "" {
		now := time.Now()
		targetDate, err = time.Parse("2006-01-02", now.Format("2006-01-02"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status" : "Error",
				"Message": "Failed to parse current date",

			})
			return
		}
	} else {
		targetDate, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status" : "Error",
				"Message": "Invalid date format. Please use YYYY-MM-DD.",
			})
			return
		}
	}

	var attendances []Models.TUAttendances
	result := Database.DB.Preload("User").Where("attendance_date = ?", targetDate).Find(&attendances)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Status" : "Not Found",
				"Message": "No attendance records found for this date.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"Status" : "Error",
			"Message": "Database query failed", 
			"Error": result.Error.Error(),
		})
		return
	}

	var response []AttendanceDetailResponse
	for _, att := range attendances {
		if att.User.UsersId != 0 {
			response = append(response, AttendanceDetailResponse{
				AttendancesUID:		   att.AttendancesUID,
				UsersUID:              att.UsersUID,
				FullName:              att.User.FullName,
				UserProfilePictureURL: att.User.ProfilePicture,
				AttendanceDate:        att.AttendanceDate,
				CheckIn:               att.CheckIn,
				StatusIn:              att.StatusIn,
				CheckInImageURL:       att.CheckInImageUrl,
				CheckOut:              att.CheckOut,
				StatusOut:             att.StatusOut,
				CheckOutImageURL:      att.CheckOutImageUrl,
			})
		}
	}

	if len(response) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"Status" : "Not Found",
			"Message": "No attendance records found for this date.", 
			"Data": []string{},
		})
		return
	}

	// 5. Kirim respons JSON
	c.JSON(http.StatusOK, gin.H{
		"Message": "Attendance data retrieved successfully",
		"Data":    response,
	})
}
