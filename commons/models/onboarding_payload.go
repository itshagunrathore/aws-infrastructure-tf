package models

type Urlstrings struct {
	enableSystem         string `default:"/dsa/components/systems/enabling/"`
	BackupSolution       string `default:"/dsa/components/backup-applications/aws-s3"`
	configureTargetGroup string `default:"/dsa/components/target-groups/s3"`
	updateMedia          string `default:"/dsa/components/mediaservers"`
	jobs                 string `default:"/dsa/jobs"`
	system               string `default:"/dsa/components/systems/teradata"`
}

type TgtPayload struct {
	AwsAccountName     string               `json:"awsAccountName"`
	IsEnabled          bool                 `json:"isEnabled"`
	Region             string               `json:"region"`
	TargetGroupName    string               `json:"targetGroupName"`
	TargetMediaBuckets []TargetMediaBuckets `json:"targetMediaBuckets"`
}
type PrefixList struct {
	PrefixName     string `json:"prefixName"`
	StorageDevices int    `json:"storageDevices"`
	PrefixId       int    `json:"prefixId,omitempt"`
}
type Buckets struct {
	BucketName string       `json:"bucketName"`
	PrefixList []PrefixList `json:"prefixList"`
	Viewpoint  bool         `json:"viewpoint,omitempt"`
}
type TargetMediaBuckets struct {
	BarMediaServer string    `json:"barMediaServer"`
	Buckets        []Buckets `json:"buckets"`
}

type DsaResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type event struct {
	DscIp          string `json:"dscIp"`
	PogNodeIp      string `json:"pogNodeIp"`
	Port           string `json:"port","default:"9090"`
	AwsAccountName string `json:"acctName,omitempty"`
	Region         string `json:"region,omitempty"`
	BucketName     string `json:"bucketName,omitempty"`
	Db_password    string `json:"dbPassword"`
	Db_user        string `json:"dbUser"`
	SystemName     string `json:"systemName"`
	RoleName       string `json:"roleName`
}

// media server struct
type MediaResponse struct {
	Validationlist interface{} `json:"validationlist"`
	Status         string      `json:"status"`
	Medias         []struct {
		ServerName string `json:"serverName"`
		Port       int    `json:"port"`
		Ips        []struct {
			IPAddress string        `json:"ipAddress"`
			Netmask   string        `json:"netmask"`
			Links     []interface{} `json:"links"`
		} `json:"ips"`
		PoolSharedPipes int           `json:"poolSharedPipes"`
		Links           []interface{} `json:"links"`
	} `json:"medias"`
	Valid bool          `json:"valid"`
	Links []interface{} `json:"links"`
}

type MediaServersConfig struct {
	IPInfo          []IPInfo `json:"ipInfo"`
	PoolSharedPipes int      `json:"poolSharedPipes"`
	Port            int      `json:"port"`
	ServerName      string   `json:"serverName"`
}
type IPInfo struct {
	IPAddress string `json:"ipAddress"`
	Netmask   string `json:"netmask"`
}

//End of media struct

// system struct
type GetSystemNames struct {
	Validationlist interface{} `json:"validationlist"`
	Status         string      `json:"status"`
	Systems        []struct {
		SystemName        string        `json:"systemName"`
		TdpID             string        `json:"tdpId"`
		IsEnabled         bool          `json:"isEnabled"`
		IrSupportSource   bool          `json:"irSupportSource"`
		IrSupportTarget   bool          `json:"irSupportTarget"`
		IrSupportOnline   bool          `json:"irSupportOnline"`
		WholeDbcSupport   bool          `json:"wholeDbcSupport"`
		IncludeDbcSupport bool          `json:"includeDbcSupport"`
		AjseSupport       bool          `json:"ajseSupport"`
		EcbbSupport       bool          `json:"ecbbSupport"`
		Links             []interface{} `json:"links"`
	} `json:"systems"`
	Valid bool          `json:"valid"`
	Links []interface{} `json:"links"`
}

type GetSystem struct {
	Validationlist interface{} `json:"validationlist"`
	Status         string      `json:"status"`
	System         struct {
		SystemName        string        `json:"systemName"`
		TdpID             string        `json:"tdpId"`
		IsEnabled         bool          `json:"isEnabled"`
		IrSupportSource   bool          `json:"irSupportSource"`
		IrSupportTarget   bool          `json:"irSupportTarget"`
		IrSupportOnline   bool          `json:"irSupportOnline"`
		WholeDbcSupport   bool          `json:"wholeDbcSupport"`
		IncludeDbcSupport bool          `json:"includeDbcSupport"`
		AjseSupport       bool          `json:"ajseSupport"`
		EcbbSupport       bool          `json:"ecbbSupport"`
		StreamsSoftlimit  int           `json:"streamsSoftlimit"`
		StreamsHardlimit  int           `json:"streamsHardlimit"`
		Links             []interface{} `json:"links"`
	} `json:"system"`
	Nodes []struct {
		NodeID       int           `json:"nodeId"`
		NodeName     string        `json:"nodeName"`
		IPAddress    []string      `json:"ipAddress"`
		NumberOfAmps int           `json:"numberOfAmps"`
		DbsNodeID    interface{}   `json:"dbsNodeId"`
		Links        []interface{} `json:"links"`
	} `json:"nodes"`
	Valid bool          `json:"valid"`
	Links []interface{} `json:"links"`
}

type DsaSystem struct {
	Password      string `json:"password"`
	SkipForceFull bool   `json:"skipForceFull"`
	// SoftLimit     int    `json:"softLimit,omitempt"`
	SystemName string `json:"systemName"`
	TdpID      string `json:"tdpId"`
	User       string `json:"user"`
}

type AwsApp struct {
	ConfigAwsRest ConfigAwsRest `json:"configAwsRest"`
}

type BucketsByRegion struct {
	Buckets          []Buckets `json:"buckets"`
	BucketsViewpoint bool      `json:"bucketsViewpoint"`
	Region           string    `json:"region"`
	Viewpoint        bool      `json:"viewpoint"`
}
type ConfigAwsRest struct {
	AcctName        string            `json:"acctName"`
	BucketsByRegion []BucketsByRegion `json:"bucketsByRegion"`
	RoleName        string            `json:"roleName"`
	Viewpoint       bool              `json:"viewpoint"`
}

// Create job struct
type CreateJob struct {
	RestJobDefinitionModel RestJobDefinitionModel `json:"restJobDefinitionModel"`
	RestJobObjectsModels   []RestJobObjectsModels `json:"restJobObjectsModels"`
	RestJobSettingsModel   RestJobSettingsModel   `json:"restJobSettingsModel"`
}
type RestJobDefinitionModel struct {
	JobName            string `json:"jobName"`
	DataDictionaryType string `json:"dataDictionaryType"`
	JobType            string `json:"jobType"`
	SourceSystem       string `json:"sourceSystem"`
	SrcUserName        string `json:"srcUserName"`
	SrcUserPassword    string `json:"srcUserPassword"`
	TargetGroupName    string `json:"targetGroupName"`
}
type RestJobObjectsModels struct {
	ParentType string `json:"parentType"`
	ObjectName string `json:"objectName"`
	ObjectType string `json:"objectType"`
	IncludeAll bool   `json:"includeAll"`
}
type RestJobSettingsModel struct {
	BlockLevelCompression string `json:"blockLevelCompression"`
	LoggingLevel          string `json:"loggingLevel"`
	Nosync                bool   `json:"nosync"`
	Nowait                bool   `json:"nowait"`
	Online                bool   `json:"online"`
	Reblock               bool   `json:"reblock"`
	SkipArchive           bool   `json:"skipArchive"`
	SkipJoinhashIndex     bool   `json:"skipJoinhashIndex"`
	SkipStats             bool   `json:"skipStats"`
	TemperatureOverride   string `json:"temperatureOverride"`
	TrackEmptyTables      bool   `json:"trackEmptyTables"`
}

// Create job struct end

type OnboardingStatus struct {
	StatusCode int
	StatusMsg  []byte
	Error      error
}
