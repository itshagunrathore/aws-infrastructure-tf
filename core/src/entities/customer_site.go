package entities

import (
	"gorm.io/datatypes"
	"time"
)

type CustomerSite struct {
	CustomerId       string
	CustomerName     string
	SiteId           string
	CustomerSiteId   int
	ServerUuid       string
	SiteTargetType   string
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
