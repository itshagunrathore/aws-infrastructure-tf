package models

type SiteTargetType string

const (
	AWS   SiteTargetType = "AWS"
	AZURE SiteTargetType = "AZURE"
	GCP   SiteTargetType = "GCP"
)
