package models

type RestJobPayload struct {
	RestJobDefinitionModel RestJobDefinitionModel `json:"restJobDefinitionModel"`
	RestJobObjectsModels   []JobObjects           `json:"restJobObjectsModels"`
	RestJobSettingsModel   JobSettings            `json:"restJobSettingsModel"`
}
type RestJobDefinitionModel struct {
	AllBackupObjects       bool    `json:"allBackupObjects,omitempty"`
	AutoRetire             bool    `json:"autoRetire,omitempty"`
	BackupName             string  `json:"backupName,omitempty"`
	BackupVersion          *int    `json:"backupVersion,omitempty"`
	DataDictionaryType     string  `json:"dataDictionaryType"`
	JobDescription         string  `json:"jobDescription"`
	JobName                string  `json:"jobName"`
	JobType                JobType `json:"jobType"`
	NextIncrementalRestore bool    `json:"nextIncrementalRestore,omitempty"`
	RetireUnits            string  `json:"retireUnits,omitempty"`
	RetireValue            int     `json:"retireValue,omitempty"`
	SaveSetAccountID       string  `json:"savesetAccountId,omitempty"`
	SaveSetPassword        string  `json:"savesetPassword,omitempty"`
	SaveSetUser            string  `json:"savesetUser,omitempty"`
	SourceSystem           string  `json:"sourceSystem"`
	SrcUserAccountID       string  `json:"srcUserAccountId,omitempty"`
	SrcUserName            string  `json:"srcUserName"`
	SrcUserPassword        string  `json:"srcUserPassword"`
	TargetGroupName        string  `json:"targetGroupName"`
	TargetSystem           string  `json:"targetSystem,omitempty"`
	TargetUserAccountID    string  `json:"targetUserAccountId,omitempty"`
	TargetUserName         string  `json:"targetUserName,omitempty"`
	TargetUserPassword     string  `json:"targetUserPassword,omitempty"`
}

type RestJobObjectsModels struct {
	ObjectName     string               `json:"objectName"`
	ObjectType     string               `json:"objectType"`
	ParentName     string               `json:"parentName"`
	ParentType     string               `json:"parentType"`
	IncludeAll     bool                 `json:"includeAll"`
	ConfigMapName  string               `json:"configMapName"`
	ExcludeObjects []RestExcludeObjects `json:"excludeObjects"`
	RenameTo       string               `json:"renameTo"`
	MapTo          string               `json:"mapTo"`
}

type RestExcludeObjects struct {
	ObjectName string `json:"objectName"`
	ObjectType string `json:"objectType"`
}

//type RestJobSettingsModel struct {
//	BlockLevelCompression     string `json:"blockLevelCompression"`
//	ConfigMapName             string `json:"configMapName"`
//	DisableFallback           bool   `json:"disableFallback"`
//	DsmainJSONLogging         bool   `json:"dsmainJsonLogging"`
//	EnableBackupForIr         bool   `json:"enableBackupForIr"`
//	EnableIncrementalRestore  bool   `json:"enableIncrementalRestore"`
//	EnableTemperatureOverride bool   `json:"enableTemperatureOverride"`
//	LoggingLevel              string `json:"loggingLevel"`
//	MapTo                     string `json:"mapTo"`
//	Nosync                    bool   `json:"nosync"`
//	Nowait                    bool   `json:"nowait"`
//	NumberParallelBuilds      int    `json:"numberParallelBuilds"`
//	Online                    bool   `json:"online"`
//}
