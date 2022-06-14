package entities

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gorm.io/datatypes"
	"time"
)

type CustomerSite struct {
	CustomerId       string
	CustomerName     string
	SiteId           string
	CustomerSiteId   int
	ServerUuid       string
	SiteTargetType   models.SiteTargetType
	SiteRegion       string
	OptInEmailReport bool
	OptInRetryFlag   bool
	IsDeleted        bool
	TimeCreated      time.Time
	TimeUpdated      time.Time
	OnboardedDate    time.Time
	SiteInfo         datatypes.JSON
	OfferingType     string
	CloudSysId       string
	CloudCustomerId  string
}

func (t *CustomerSite) TableName() string {
	return "customer_site"
}
