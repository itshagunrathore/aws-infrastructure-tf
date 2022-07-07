package dsa

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"io/ioutil"
	"net/http"
)

type dsaDriver struct {
	host         string
	port         string
	userName     string
	password     string
	secure       bool
	clientId     string
	clientSecret string
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
	fmt.Println(endpoint)
	fmt.Println(siteTargetType)
	url := d.getBaseUrl() + endpoint
	resp, err := GetDsa(url)

	if err != nil {
		return models.TargetGroupsResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return models.TargetGroupsResponse{}, errors.New("test")
	}

	var TargetGroupsResponse models.TargetGroupsResponse
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &TargetGroupsResponse)

	if err != nil {
		return models.TargetGroupsResponse{}, err
	}
	// check for valid response
	return TargetGroupsResponse, nil

}

func (d *dsaDriver) GetSystemNames() (models.SystemsResponse, error) {

	url := d.getBaseUrl() + "/dsa/components/systems/teradata"
	resp, err := GetDsa(url)
	if err != nil {
		return models.SystemsResponse{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return models.SystemsResponse{}, err
	}
	var systemsResponse models.SystemsResponse

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &systemsResponse)

	if err != nil {
		return models.SystemsResponse{}, err
	}

	return systemsResponse, nil
}

func (d *dsaDriver) PostJob(restJobPayload models.RestJobPayload) error {
	url := d.getBaseUrl() + "/dsa/jobs"
	jsonString, err := json.Marshal(restJobPayload)
	if err != nil {
		return err
	}
	resp, err := PostDsa(url, jsonString)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to create job on dsa")
	}
	var createJobResponse map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &createJobResponse)
	if err != nil {
		return err
	}
	//if createJobResponse[""]
	return errors.New("test")
}

func (d *dsaDriver) getBaseUrl() string {
	return "https://" + d.host + ":" + string(d.port)
}
