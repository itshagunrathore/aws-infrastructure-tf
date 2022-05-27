package entities

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"time"
)

type LatestJobSession struct {
	JobSessionId         int `gorm:"primary_key;not null;autoIncrement"`
	JobId                int `gorom`
	JobStartTime         *time.Time
	JobEndTime           *time.Time
	LatestStatus         string
	LatestStatusCode     string
	LatestStatusDetails  string
	latestDsaJobStatus   string
	RetryCount           int
	ServerExecutionId    *string
	BackupSetSizeInBytes int64
	UpdatedAt            time.Time `gorm:"column:time_updated;not null"`
	CreatedAt            time.Time `gorm:"column:time_created;not null"`
	RunType              *models.RunType
}
