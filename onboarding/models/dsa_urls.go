package models

type DsaUrlString string

const (
	ConfigSystem   DsaUrlString = "/dsa/components/systems/teradata"
	GetSystems     DsaUrlString = ""
	ConfigAwsTGT   DsaUrlString = "/dsa/components/target-groups/s3"
	ConfigAzureTGT DsaUrlString = ""
	ConfigGcpTGT   DsaUrlString = ""
	GetTGT         DsaUrlString = ""
	MediaServer    DsaUrlString = "/dsa/components/mediaservers"
	S3App          DsaUrlString = "/dsa/components/backup-applications/aws-s3"
	AzureApp       DsaUrlString = ""
	GcpApp         DsaUrlString = ""
	ConfigureJob   DsaUrlString = "/dsa/jobs"
)
