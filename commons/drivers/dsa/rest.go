package dsa

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"net/http"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func GetDsa(url string) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PostDsa(url string, jsonString []byte) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonString))

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func PostConfigDsc(url string, payload interface{}, DsaStatus *models.DetailedStatus) ([]byte, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	data, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		DsaStatus.StatusCode = 503
		DsaStatus.StepResponse = "Failed to create http request"
		DsaStatus.Error = err
		return []byte(DsaStatus.StepResponse), err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		DsaStatus.Error = err
		DsaStatus.StatusCode = 502
		DsaStatus.StepResponse = "DSA Gateway Timeout"
		return []byte("Failed to invoke dsa api"), err
	}
	if resp.StatusCode != 200 {
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		DsaStatus.Error = err
		DsaStatus.StatusCode = resp.StatusCode
		DsaStatus.StepResponse = resp.Status
		return []byte("failed"), err
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		DsaStatus.Error = err
		DsaStatus.StatusCode = 400
		DsaStatus.StepResponse = string(response)
		log.Fatalln(err)
	}
	DsaStatus.Error = err
	DsaStatus.StatusCode = resp.StatusCode
	DsaStatus.StepResponse = resp.Status
	return response, err
}

func GetConfigDsc(url string, DsaStatus *models.DetailedStatus) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		DsaStatus.StatusCode = 503
		DsaStatus.StepResponse = "Failed to create http request"
		DsaStatus.Error = err
		return []byte(DsaStatus.StepResponse), err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		DsaStatus.Error = err
		DsaStatus.StatusCode = 502
		DsaStatus.StepResponse = "DSA Gateway Timeout"
		return []byte("Failed to invoke dsa api"), err
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		DsaStatus.Error = err
		DsaStatus.StatusCode = 400
		DsaStatus.StepResponse = string(response)
		return response, err
	}
	if resp.StatusCode != 200 {
		DsaStatus.Error = err
		DsaStatus.StatusCode = resp.StatusCode
		DsaStatus.StepResponse = resp.Status
		return response, err
	}
	DsaStatus.Error = err
	DsaStatus.StatusCode = resp.StatusCode
	DsaStatus.StepResponse = resp.Status

	return response, err
}
