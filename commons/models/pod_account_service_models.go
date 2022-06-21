package models

import "time"

type DscInstanceDetails struct {
	ClientName   string `json:"clientName"`
	DsaPrivateIP string `json:"dsaPrivateIp,omitempty"`
	DsaPublicIP  string `json:"dsaPublicIp,omitempty"`
	Status       string `json:"status"`
	Error        string `json:"error,omitempty"`
}

type AccountDetails struct {
	XItemType     string `json:"X_ITEM_TYPE"`
	AccountID     string `json:"accountId"`
	AccountName   string `json:"accountName"`
	AccountStatus string `json:"accountStatus"`
	AccountType   string `json:"accountType"`
	Active        bool   `json:"active"`
	Components    []struct {
		AccountID string `json:"accountId"`
		Active    bool   `json:"active"`
		Compute   []struct {
			NodeCount int    `json:"nodeCount"`
			NodeType  string `json:"nodeType"`
		} `json:"compute"`
		CreatedAt      time.Time `json:"createdAt"`
		ID             string    `json:"id"`
		Infrastructure struct {
			Idp struct {
			} `json:"IDP"`
			CapacityReservation struct {
			} `json:"capacityReservation"`
			DataVolumeType struct {
			} `json:"dataVolumeType"`
			DefaultDataVolumeType struct {
				Iops       string `json:"IOPS"`
				Throughput string `json:"throughput"`
				VolumeType string `json:"volumeType"`
			} `json:"defaultDataVolumeType"`
			DefaultImageID string `json:"defaultImageId"`
			IamProfileArn  string `json:"iamProfileArn"`
			ImageID        string `json:"imageId"`
			ImageName      string `json:"imageName"`
			KeyName        string `json:"keyName"`
			MssUser        struct {
				AccessKey string `json:"accessKey"`
				PolicyArn string `json:"policyArn"`
				SecretKey string `json:"secretKey"`
				UserName  string `json:"userName"`
			} `json:"mssUser"`
			NosfsUser struct {
				AccessKey string `json:"accessKey"`
				PolicyArn string `json:"policyArn"`
				SecretKey string `json:"secretKey"`
				UserName  string `json:"userName"`
			} `json:"nosfsUser"`
			OverProvision   bool   `json:"overProvision"`
			PlacementGroup  string `json:"placementGroup"`
			PodAddress      string `json:"podAddress"`
			PodId           string `json:"podId"`
			PogrouterSubnet struct {
			} `json:"pogrouterSubnet"`
			PrivateSubnet1     string `json:"privateSubnet1"`
			PrivateSubnet1Cidr string `json:"privateSubnet1Cidr"`
			PrivateSubnet2     struct {
			} `json:"privateSubnet2"`
			PrivateSubnet2Cidr struct {
			} `json:"privateSubnet2Cidr"`
			PublicIP          bool   `json:"publicIp"`
			PublicSubnet1Cidr string `json:"publicSubnet1Cidr"`
			PublicSubnet2Cidr struct {
			} `json:"publicSubnet2Cidr"`
			S3Buckets struct {
				Mss struct {
					BucketName      string `json:"bucketName"`
					EncryptionArn   string `json:"encryptionArn"`
					EncryptionKeyID string `json:"encryptionKeyId"`
				} `json:"mss"`
				Nosfs struct {
					BucketName      string `json:"bucketName"`
					EncryptionArn   string `json:"encryptionArn"`
					EncryptionKeyID string `json:"encryptionKeyId"`
				} `json:"nosfs"`
				Support struct {
					BucketName      string `json:"bucketName"`
					EncryptionArn   string `json:"encryptionArn"`
					EncryptionKeyID string `json:"encryptionKeyId"`
				} `json:"support"`
			} `json:"s3buckets"`
			Stage    string `json:"stage"`
			SubnetID string `json:"subnetId"`
		} `json:"infrastructure"`
		Job     interface{}   `json:"job"`
		Name    string        `json:"name"`
		Network []interface{} `json:"network"`
		Nodes   []struct {
			AccountID        string `json:"accountId"`
			Active           bool   `json:"active"`
			AvailabilityZone string `json:"availabilityZone"`
			ComponentID      string `json:"componentId"`
			Compute          struct {
				InstanceName string `json:"instanceName"`
				InstanceType string `json:"instanceType"`
			} `json:"compute"`
			ID             string `json:"id"`
			Infrastructure struct {
				Eip                 bool `json:"EIP"`
				CapacityReservation struct {
				} `json:"capacityReservation"`
				DataVolumeType struct {
				} `json:"dataVolumeType"`
				DefaultDataVolumeType struct {
					Iops       string `json:"IOPS"`
					Throughput string `json:"throughput"`
					VolumeType string `json:"volumeType"`
				} `json:"defaultDataVolumeType"`
				DefaultImageID string `json:"defaultImageId"`
				IamProfileArn  string `json:"iamProfileArn"`
				ImageID        string `json:"imageId"`
				KeyName        string `json:"keyName"`
				OverProvision  bool   `json:"overProvision"`
				PlacementGroup string `json:"placementGroup"`
				PodAddress     string `json:"podAddress"`
				PodId          string `json:"podId"`
				PublicIP       bool   `json:"publicIp"`
				Stage          string `json:"stage"`
				SubnetID       string `json:"subnetId"`
				UserData       string `json:"userData"`
			} `json:"infrastructure"`
			InstanceID string `json:"instanceId"`
			Name       string `json:"name"`
			Network    struct {
				PrivateDNS       string   `json:"privateDns"`
				PrivateIP        string   `json:"privateIp"`
				PublicDNS        string   `json:"publicDns"`
				PublicIP         string   `json:"publicIp"`
				SecondaryIPCount int      `json:"secondaryIpCount"`
				SecondaryIps     []string `json:"secondaryIps"`
			} `json:"network"`
			NodeID         string   `json:"nodeId"`
			NodeType       string   `json:"nodeType"`
			Platform       string   `json:"platform"`
			PmaID          int      `json:"pmaId"`
			PrivateIP      string   `json:"privateIp"`
			PublicIP       string   `json:"publicIp"`
			Region         string   `json:"region"`
			SecondaryIps   []string `json:"secondaryIps"`
			ServiceAddress string   `json:"serviceAddress"`
			Storage        struct {
				ArrayType   string `json:"arrayType"`
				DataDevices []struct {
					DiskType   string `json:"diskType"`
					Name       string `json:"name"`
					Partitions []struct {
						End       int    `json:"end"`
						Name      string `json:"name"`
						Number    int    `json:"number"`
						PdiskID   int    `json:"pdiskId"`
						PdiskLink string `json:"pdiskLink"`
						Sectors   int    `json:"sectors"`
						Start     int    `json:"start"`
						Usage     string `json:"usage"`
					} `json:"partitions"`
					Sectors     int    `json:"sectors"`
					VDeviceName string `json:"vDeviceName"`
					VDiskID     string `json:"vDiskId"`
				} `json:"dataDevices"`
				Encryption bool `json:"encryption"`
				Ephemeral  struct {
					DiskType    string `json:"diskType"`
					NamePattern string `json:"namePattern"`
					Usage       struct {
						Noscache int `json:"noscache"`
					} `json:"usage"`
				} `json:"ephemeral"`
				MaxPersistentDataGB int `json:"maxPersistentDataGB"`
				MinPersistentDataGB int `json:"minPersistentDataGB"`
				PersistentDataGB    int `json:"persistentDataGB"`
				RootDevice          struct {
					DiskType    string `json:"diskType"`
					Name        string `json:"name"`
					VDeviceName string `json:"vDeviceName"`
					VDiskID     string `json:"vDiskId"`
				} `json:"rootDevice"`
				RootGB      int `json:"rootGB"`
				VolumeCount int `json:"volumeCount"`
			} `json:"storage"`
			SystemID   string `json:"systemId"`
			SystemName string `json:"systemName"`
		} `json:"nodes"`
		Platform string `json:"platform"`
		Region   string `json:"region"`
		State    struct {
			Name            string    `json:"name"`
			ObjID           string    `json:"objId"`
			RunningAt       time.Time `json:"runningAt"`
			RunningAtLocal  string    `json:"runningAtLocal"`
			SystemID        string    `json:"systemId"`
			ValidOperations []string  `json:"validOperations"`
		} `json:"state"`
		Status struct {
			Name     string `json:"name"`
			ObjID    string `json:"objId"`
			SystemID string `json:"systemId"`
		} `json:"status"`
		Storage []struct {
			MaxPersistentDataGB int    `json:"maxPersistentDataGB"`
			NodeType            string `json:"nodeType"`
			PersistentDataGB    int    `json:"persistentDataGB"`
		} `json:"storage"`
		SystemID   string `json:"systemId"`
		SystemName string `json:"systemName"`
		Type       string `json:"type"`
	} `json:"components"`
	CreatedAt      time.Time `json:"createdAt"`
	DbcPassword    string    `json:"dbcPassword"`
	DeployedAt     time.Time `json:"deployedAt"`
	ExternalSiteID string    `json:"externalSiteId"`
	Infrastructure struct {
		Eip                   bool   `json:"EIP"`
		AccountID             string `json:"accountId"`
		DefaultIamProfileArn  string `json:"defaultIamProfileArn"`
		DefaultImageID        string `json:"defaultImageId"`
		DefaultRegion         string `json:"defaultRegion"`
		DefaultSecurityGroups string `json:"defaultSecurityGroups"`
		Iam                   struct {
			Tpa struct {
				InstanceProfile string `json:"instanceProfile"`
				PolicyArn       string `json:"policyArn"`
				RoleArn         string `json:"roleArn"`
				RoleName        string `json:"roleName"`
			} `json:"tpa"`
		} `json:"iam"`
		KeyName            string `json:"keyName"`
		PemKey             string `json:"pemKey"`
		PodAddress         string `json:"podAddress"`
		PodId              string `json:"podid"`
		PrivateSubnet1     string `json:"privateSubnet1"`
		PrivateSubnet1Cidr string `json:"privateSubnet1Cidr"`
		PublicIP           bool   `json:"publicIp"`
		PublicSubnet1Cidr  string `json:"publicSubnet1Cidr"`
		S3Buckets          struct {
			Mss struct {
				BucketName      string `json:"bucketName"`
				EncryptionArn   string `json:"encryptionArn"`
				EncryptionKeyID string `json:"encryptionKeyId"`
			} `json:"mss"`
			Nosfs struct {
				BucketName      string `json:"bucketName"`
				EncryptionArn   string `json:"encryptionArn"`
				EncryptionKeyID string `json:"encryptionKeyId"`
			} `json:"nosfs"`
			Support struct {
				BucketName      string `json:"bucketName"`
				EncryptionArn   string `json:"encryptionArn"`
				EncryptionKeyID string `json:"encryptionKeyId"`
			} `json:"support"`
		} `json:"s3buckets"`
		Stage               string `json:"stage"`
		SystemID            string `json:"systemId"`
		TenantSecurityGroup string `json:"tenantSecurityGroup"`
		Vpc                 string `json:"vpc"`
	} `json:"infrastructure"`
	Name              string `json:"name"`
	Platform          string `json:"platform"`
	PodTenantAddress  string `json:"podTenantAddress"`
	PogRouterSystemID string `json:"pogRouterSystemId"`
	Region            string `json:"region"`
	RegionDetails     struct {
		Endpoint string `json:"endpoint"`
		Name     string `json:"name"`
		Type     string `json:"type"`
	} `json:"regionDetails"`
	Release    interface{} `json:"release"`
	Stage      string      `json:"stage"`
	SystemName string      `json:"systemName"`
}
