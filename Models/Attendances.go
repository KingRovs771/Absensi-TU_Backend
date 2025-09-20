package Models

type TUAttendances struct {
	AttendancesId    int    `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"attendances_id"`
	AttendancesUID   string `gorm:"type:varchar" json:"attendances_uid"`
	AttendanceDate   string `gorm:"type:date" json:"attendance_date"`
	CheckIn          string `gorm:"type:time" json:"check_in"`
	StatusIn         string `gorm:"type:varchar" json:"status_in"`
	CheckInImageUrl  string `gorm:"type:varchar" json:"check_in_image_url"`
	CheckOut         string `gorm:"type:time" json:"check_out"`
	StatusOut        string `gorm:"type:varchar" json:"status_out"`
	CheckOutImageUrl string `gorm:"type:varchar" json:"check_out_image_url"`
}
