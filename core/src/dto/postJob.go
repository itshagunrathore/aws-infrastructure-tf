package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type PostJobDto struct {
	Name                string                 `json:"name"`
	Description         string                 `json:"description"`
	JobType             models.JobType         `json:"job_type"`
	IsActive            bool                   `json:"is_active,omitempty"`
	NoOfRetentionCopies int                    `json:"no_of_retention_copies,omitempty"`
	IsAutoAbortActive   bool                   `json:"is_auto_abort_active,omitempty"`
	AutoAbortInMinutes  int                    `json:"auto_abort_in_minutes,omitempty"`
	BackupMechanism     models.BackupMechanism `json:"backup_mechanism"`
	BackupType          models.BackupType      `json:"backup_type"`
	DsaJobDefinition    DsaJobDefinition       `json:"dsa_job_definition,omitempty"`
}

type DsaJobDefinition struct {
	JobObjects  []models.JobObjects `json:"job_objects,omitempty"`
	JobSettings models.JobSettings  `json:"job_settings,omitempty"`
}

func (dto PostJobDto) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required),
		validation.Field(&dto.Description, validation.Required),
		validation.Field(&dto.JobType, validation.Required, validation.In(models.Backup, models.Restore)),
		validation.Field(&dto.IsActive, validation.Required),
		validation.Field(&dto.BackupMechanism, validation.Required, validation.In(models.DSA, models.CDP)),
		validation.Field(&dto.DsaJobDefinition),
		//validation.Required.When(dto.DsaJobDefinition != nil), validation.Nil.When(dto.DsaJobDefinition == nil)
	)
}

func (dsaJobDefinition DsaJobDefinition) Validate() error {
	return validation.ValidateStruct(&dsaJobDefinition,
		validation.Field(&dsaJobDefinition.JobObjects),
	)
}
