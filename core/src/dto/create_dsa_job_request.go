package dto

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"

type CreateDsaJobRequest struct {
	JobName        string
	Description    string
	JobType        models.JobType
	accountId      string
	siteTargetType string
	BaaSJobId      int
	JobSettings    models.JobSettings
	JobObjects     []models.JobObjects
}
