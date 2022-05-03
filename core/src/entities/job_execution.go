package entities

import "time"

type JobExecutionEntity struct {
	RestoreExecutionId int       `gorm:"column:restore_execution_id;not null"`
	JobExecutionId     int       `gorm:"column:job_execution_id"`
	StartTime          time.Time `gorm:"column:start_time"`
	EndTime            time.Time `gorm:"column:end_time"`
	Status             string    `gorm:"column:status"`
	StatusCode         int       `gorm:"column:status_code"`
	StatusDetails      string    `gorm:"column:status_details"`
	Progress           string    `gorm:"column:progress"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
	TicketId           string    `gorm:"column:ticket_id"`
	NotificationId     string    `gorm:"column:notification_id"`
	SiteId             string    `gorm:"column:site_id;not null"`
}
