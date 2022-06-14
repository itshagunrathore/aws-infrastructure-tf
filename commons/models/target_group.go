package models

type TargetGroupsResponse struct {
	Validationlist Validationlist `json:"validationlist"`
	Status         string         `json:"status"`
	S3Target       []S3Target     `json:"s3Target"`
	Azure          []AzureTarget  `json:"azure"`
	Gcp            []GcpTarget    `json:"gcp"`
	Valid          bool           `json:"valid"`
}

type S3Target struct {
	TargetGroupName string `json:"targetGroupName"`
}

type AzureTarget struct {
	TargetGroupName string `json:"targetGroupName"`
}

type GcpTarget struct {
	TargetGroupsName string `json:"targetGroupName"`
}
