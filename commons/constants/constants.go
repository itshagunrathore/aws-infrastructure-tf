package constants

const (
	ClientName         string = "baas" //
	Running            string = "running"
	Deploying          string = "deploying"
	Terminating        string = "terminating"
	ProvisionDsaPath   string = "/v1/accounts/{accountId}/dsa"
	DeprovisionDsaPath string = "/v1/accounts/{accountId}/{clientName}/{clientSessionId}/dsa"
)
