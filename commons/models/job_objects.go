package models

type JobObjects struct {
	ConfigMapName  string           `json:"configMapName,omitempty"`
	ExcludeObjects []ExcludeObjects `json:"excludeObjects,omitempty"`
	IncludeAll     bool             `json:"includeAll"`
	MapTo          string           `json:"mapTo,omitempty"`
	ObjectName     string           `json:"objectName"`
	ObjectType     string           `json:"objectType"`
	ParentName     string           `json:"parentName"`
	ParentType     string           `json:"parentType"`
	RenameTo       string           `json:"renameTo,omitempty"`
}

func NewJobObjects(configMapName string, excludeObjects []ExcludeObjects, includeAll bool, mapTo string, objectName string, objectType string, parentName string, parentType string, renameTo string) *JobObjects {
	return &JobObjects{ConfigMapName: configMapName, ExcludeObjects: excludeObjects, IncludeAll: includeAll, MapTo: mapTo, ObjectName: objectName, ObjectType: objectType, ParentName: parentName, ParentType: parentType, RenameTo: renameTo}
}
