package dto

import (
	"time"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type GetJobDto struct {
	JobID                int                    `json:"job_id"`
	Name                 string                 `json:"name"`
	Description          string                 `json:"description"`
	SiteID               string                 `json:"site_id"`
	IsActive             bool                   `json:"is_active"`
	Priority             int                    `json:"priority"`
	JobType              models.JobType         `json:"job_type"`
	SiteTargetType       models.SiteTargetType  `json:"site_target_type"`
	LastExecutionDetails LastExecutionDetails   `json:"last_execution_details"`
	NoOfRetentionCopies  int                    `json:"no_of_retention_copies"`
	IsAutoAbortActive    bool                   `json:"is_auto_abort_active"`
	AutoAbortInMinutes   int                    `json:"auto_abort_in_minutes"`
	NextRunTime          string                 `json:"next_run_time"`
	BackupMechanism      models.BackupMechanism `json:"backup_mechanism"`
	DsaJobDefinition     DsaJobDefinition       `json:"dsa_job_definition"`
}
type LastExecutionDetails struct {
	Status        string     `json:"status"`
	BackupSetSize int64      `json:"backup_set_size"`
	StartTime     *time.Time `json:"start_time"`
	EndTime       *time.Time `json:"end_time"`
}
