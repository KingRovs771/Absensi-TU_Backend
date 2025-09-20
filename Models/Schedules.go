package Models

type TUSchedules struct {
	SchedulesId int    `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"schedules_id"`
	ScheduleUID string `gorm:"type:varchar" json:"schedule_uid"`
	StartTime   string `gorm:"type:time" json:"start_time"`
	EndTime     string `gorm:"type:time" json:"end_time"`
	CreateAt    int64  `gorm:"autoCreateTime" json:"create_at"`
}
