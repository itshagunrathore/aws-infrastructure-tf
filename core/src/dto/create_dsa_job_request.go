package dto

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"

type CreateDsaJobRequest struct {
	JobName        string
	Description    string
	JobType        models.JobType
	AccountId      string
	SiteTargetType models.SiteTargetType
	BaaSJobId      int
	JobSettings    models.JobSettings
	JobObjects     []models.JobObjects
}
