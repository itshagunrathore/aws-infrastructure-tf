package entities

import (
	"time"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type Tenant struct {
	CustomerName     string
	SiteId           string
	TenantId         string
	PodId            string
	CustomerSiteId   int    `gorm:"primary_key;not null;autoIncrement"`
	ServerUuid       string `gorm:"type:uuid"`
	SiteTargetType   models.SiteTargetType
	SiteRegion       string
	OptInEmailReport bool `gorm:"default:false"`
	OptInRetryFlag   bool `gorm:"default:false"`
	IsDeleted        bool `gorm:"default:false"`
	TimeCreated      time.Time
	TimeUpdated      time.Time
	OnboardedDate    time.Time
	OfferingType     string
}

func (t *Tenant) TableName() string {
	return "tenant"
}