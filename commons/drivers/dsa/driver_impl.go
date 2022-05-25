package dsa

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/customerrors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"io/ioutil"
	"net/http"
)

type dsaDriver struct {
	host          string
	port          string
	userName      string
	password      string
	certSignature string
	secure        bool
	clientId      string
	clientSecret  string
}

func NewDsaDriver(host string, port string, userName string, password string, clientId string, clientSecret string) DsaDriver {
	return &dsaDriver{
		host:         host,
		port:         port,
		userName:     userName,
		password:     password,
		secure:       true,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (d *dsaDriver) GetTargetGroup(siteTargetType models.SiteTargetType) (models.TargetGroupsResponse, error) {
	endpoint := ""
	if siteTargetType == models.AWS {
		endpoint = "/dsa/components/target-groups/s3"
	} else if siteTargetType == models.AZURE {
		endpoint = "/dsa/components/target-groups/azure"
	} else if siteTargetType == models.GCP {
		endpoint = "/dsa/components/target-groups/gcp"
	}

	url := d.getBaseUrl() + endpoint
	resp, err := GetDsa(url)

	if err != nil {
		return models.TargetGroupsResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return models.TargetGroupsResponse{}, errors.New("test")
	}

	var TargetGroupsReponse models.TargetGroupsResponse
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &TargetGroupsReponse)

	if err != nil {
		return models.TargetGroupsResponse{}, err
	}
	// check for valid response
	return TargetGroupsReponse, nil

}

func (d *dsaDriver) SystemNames() (models.SystemsResponse, error) {

	url := d.getBaseUrl() + "/dsa/components/systems/teradata"
	resp, err := GetDsa(url)
	if err != nil {
		return models.SystemsResponse{}, customerrors.DsaError{}
	}
	if resp.StatusCode != http.StatusOK {
		return models.SystemsResponse{}, customerrors.DsaError{}
	}
	var systemsResponse models.SystemsResponse

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &systemsResponse)

	if err != nil {
		return models.SystemsResponse{}, customerrors.DsaError{}
	}

	return systemsResponse, nil
}

func (d *dsaDriver) PostJob(model models.RestJobPayload) error {
	url := d.getBaseUrl() + "/dsa/jobs"
	jsonString, err := json.Marshal(model)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonString))
	resp, err := PostDsa(url, jsonString)

	fmt.Println(resp)
	return errors.New("test")
}

func (d dsaDriver) getBaseUrl() string {
	return "https://" + d.host + ":" + string(d.port)
}

func (d dsaDriver) processResponse(resp *http.Response) {

	var dsaResponse map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &dsaResponse)
	if err != nil {
		return
	}
	if !dsaResponse["valid"].(bool) {

	}

}
