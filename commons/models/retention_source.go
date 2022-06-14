package models

type RetentionSource string

const (
	S3           RetentionSource = "S3"
	Blob         RetentionSource = "BLOB"
	CloudStorage RetentionSource = "CloudStorage"
	DataDomain   RetentionSource = "DATA_DOMAIN"
)
