package constants

const (
	ClientName         = "BaaSClient" //
	Running            = "running"
	Deploying          = "deploying"
	Terminating        = "terminating"
	ProvisionDsaPath   = "/v1/accounts/{accountId}/dsa"
	DeprovisionDsaPath = "/v1/accounts/{accountId}/{clientName}/{clientSessionId}/dsa"
)
