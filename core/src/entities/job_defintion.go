package entities

import "time"

type retentionSource string

const (
	S3           retentionSource = "S3"
	Blob         retentionSource = "BLOB"
	CloudStorage retentionSource = "CloudStorage"
	DataDomain   retentionSource = "DATA_DOMAIN"
)

type backupMechanism string

const (
	DSA backupMechanism = "DSA"
	CDP backupMechanism = "CDP"
)

type jobType string

const (
	Restore     jobType = "RESTORE"
	Replication jobType = "REPLICATION"
	Backup      jobType = "BACKUP"
)

type backupType string

const (
	Full  backupType = "FULL"
	Delta backupType = "DELTA"
)

type JobDefinition struct {
	JobId                int       `gorm:"primary_key;not null"`
	Name                 string    `gorm:"column:job_name"`
	UpdatedAt            time.Time `gorm:"column:time_updated"`
	CreatedAt            time.Time `gorm:"column:time_created"`
	Description          string
	Status               string
	StatusCode           string
	StatusDetails        string
	IsDeleted            bool
	IsHidden             bool
	BackupSegment        string
	DataPhase            string
	CustomerSiteId       int
	RetentionSource      retentionSource `gorm:"not null"`
	IsReplicationEnabled bool            `gorm:"not null"`
	IsManaged            bool
	DeltaCount           int
	BackupMechanism      backupMechanism
	SourceSite           string
	TargetSite           string
	TargetInfoId         int
	IsActive             bool
	JobPriority          int     `gorm:"not null"`
	JobType              jobType `gorm:"not null"`
	BackupType           backupType
	RetentionCopiesCount int
	IsAutoAbortActive    bool         `gorm:"not null"`
	AutoAbortInMin       int          `gorm:"not null"`
	CustomerSite         CustomerSite `gorm:"foreignKey:CustomerSiteId"`
}

func (t *JobDefinition) TableName() string {
	return "job_definition"
}
