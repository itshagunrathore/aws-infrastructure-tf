package httpclient

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"time"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
)

type HttpClient interface {
	Get(url string) ([]byte, int, error)
	Post(url string, request bytes.Buffer) ([]byte, int, error)
	Delete(url string, request ...bytes.Buffer) ([]byte, int, error)
}
type httpClient struct {
	client *http.Client
}

func NewHttpClient(secure bool) HttpClient {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	if !secure {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
	}
	return &httpClient{client: client}
}

// setup other required headers, auth when required
func setupHeaders(req *http.Request) {
	req.Header.Add("Accept", `application/json`)

	if req.Method == http.MethodPost {
		req.Header.Add("Content-Type", `application/json`)
	}
}

func (h *httpClient) Get(url string) ([]byte, int, error) {
	log.Infof("Recieved GET request for endpoint: %s", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	setupHeaders(req)

	resp, err := h.client.Do(req)
	log.Infof("Response for GET request: %v", resp.StatusCode)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func (h *httpClient) Post(url string, request bytes.Buffer) ([]byte, int, error) {
	log.Infof("Recieved POST request for endpoint: %s", url)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(request.Bytes()))
	if err != nil {
		return nil, 0, err
	}
	setupHeaders(req)
	resp, err := h.client.Do(req)
	log.Infof("Response for POST request: %v", resp.StatusCode)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, resp.StatusCode, nil
}
func (h *httpClient) Delete(url string, request ...bytes.Buffer) ([]byte, int, error) {
	var req *http.Request
	var err error
	log.Infof("Recieved DELETE request for endpoint: %s", url)
	req, err = h.GetDeleteRequest(url, request...)
	if err != nil {
		return nil, 0, err
	}
	setupHeaders(req)
	resp, err := h.client.Do(req)
	log.Infof("Response for DELETE request: %v", resp.StatusCode)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
func (h *httpClient) GetDeleteRequest(url string, request ...bytes.Buffer) (*http.Request, error) {
	if len(request) == 0 {
		return http.NewRequest(http.MethodDelete, url, nil)
	}
	r := request[0]
	return http.NewRequest(http.MethodDelete, url, bytes.NewReader(r.Bytes()))
}
