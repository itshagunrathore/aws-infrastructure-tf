package dsa

import "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"

type DsaDriver interface {
	GetTargetGroup(siteTargetType models.SiteTargetType) (models.TargetGroupsResponse, error)
	GetSystemNames() (models.SystemsResponse, error)
	PostJob(payload models.RestJobPayload) error
}
