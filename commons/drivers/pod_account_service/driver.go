package podaccountservice

import (
	"bytes"
	"encoding/json"
	"net/http"

	httpClient "gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/http"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

type PodAccountService interface {
	ProvisionDsa(url, path string, secure bool, provisionDsaModel models.ProvisionDsaModel) (*http.Response, error)
	GetDsaStatus(url, path string, secure bool) (*http.Response, error)
}

type podAccountService struct {
}

func NewPodAccountService() *podAccountService {
	return &podAccountService{}
}
func (p *podAccountService) ProvisionDsa(url, path string, secure bool, provisionDsaModel models.ProvisionDsaModel) (*http.Response, error) {
	httpClient := httpClient.NewHttpClient()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(provisionDsaModel)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Post(url+path, false, buf)
	if err != nil {
		return nil, err
	}
	log.Info(resp)
	return resp, err
}
func (p *podAccountService) GetDsaStatus(url, path string, secure bool) (*http.Response, error) {
	httpClient := httpClient.NewHttpClient()
	resp, err := httpClient.Get(url+path, secure)
	if err != nil {
		return nil, err
	}
	log.Info(resp)
	return resp, err
}
func (p *podAccountService) DeprovisionDsa(url, path string, secure bool) (*http.Response, error) {
	httpClient := httpClient.NewHttpClient()
	resp, err := httpClient.Delete(url+path, secure)
	if err != nil {
		return nil, err
	}
	log.Info(resp)
	return resp, err
}
