package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

type JobObjects struct {
	ObjectName     string           `json:"object_name"`
	ObjectType     string           `json:"object_type"`
	ParentName     string           `json:"parent_name"`
	ParentType     string           `json:"parent_type"`
	IncludeAll     bool             `json:"include_all"`
	ConfigMapName  string           `json:"config_map_name,omitempty"`
	ExcludeObjects []ExcludeObjects `json:"exclude_objects,omitempty"`
	RenameTo       string           `json:"renameTo,omitempty"`
	MapTo          string           `json:"mapTo,omitempty"`
}

func (jobObjects JobObjects) Validate() error {
	return validation.ValidateStruct(&jobObjects,
		validation.Field(&jobObjects.ObjectName, validation.Required),
		validation.Field(&jobObjects.ObjectType, validation.Required),
		validation.Field(&jobObjects.ParentName, validation.Required),
		validation.Field(&jobObjects.ParentType, validation.Required),
		validation.Field(&jobObjects.IncludeAll, validation.Required),
	)
}
