package entities

import (
	"gorm.io/datatypes"
	"time"
)

type CustomerSite struct {
	customerId       string
	customerName     string
	siteId           string
	customerSiteId   int
	serverUuid       string
	siteTargetType   string
	siteRegion       string
	optInEmailReport bool
	optInRetryFlag   bool
	isDeleted        bool
	timeCreated      time.Time
	timeUpdated      time.Time
	onboardedDate    time.Time
	siteInfo         datatypes.JSON
	offeringType     string
	cloudSysId       string
	cloudCustomerId  string
}
