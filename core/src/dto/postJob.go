package dto

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type PostJobDto struct {
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	JobType             models.JobType         `json:"job_type"`
	IsActive            bool                   `json:"is_active"`
	NoOfRetentionCopies *int                   `json:"no_of_retention_copies"`
	IsAutoAbortActive   bool                   `json:"is_auto_abort_active"`
	AutoAbortInMinutes  *int                   `json:"auto_abort_in_minutes"`
	BackupMechanism     models.BackupMechanism `json:"backup_mechanism"`
	BackupType          models.BackupType      `json:"backup_type"`
	DsaJobDefinition    DsaJobDefinition       `json:"dsa_job_definition,omitempty"`
}

type DsaJobDefinition struct {
	JobObjects  []models.JobObjects `json:"job_objects,omitempty"`
	JobSettings models.JobSettings  `json:"job_settings,omitempty"`
}
