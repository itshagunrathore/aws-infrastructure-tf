package dsa

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"

type DsaDriver interface {
	GetTargetGroup(siteTargetType models.SiteTargetType) (models.TargetGroupsResponse, error)
	SystemNames() (models.SystemsResponse, error)
	PostJob(paylaod models.RestJobPayload) error
}
