package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TUSchedules struct {
	SchedulesId int    `gorm:"primary_key;AUTO_INCREMENT;uniqueIndex" json:"schedules_id"`
	ScheduleUID string `gorm:"type:varchar" json:"schedule_uid"`
	Day         string `gorm:"type:varchar" json:"day"`
	StartTime   string `gorm:"type:time" json:"start_time"`
	EndTime     string `gorm:"type:time" json:"end_time"`
	CreateAt    int64  `gorm:"autoCreateTime" json:"create_at"`
}

func (schedules *TUSchedules) BeforeSaveSchedules(*gorm.DB) error{
	UniqueID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	schedules.ScheduleUID = UniqueID.String()

	return nil
}