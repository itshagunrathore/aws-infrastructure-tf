package httpClient

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

type HttpClient interface {
	Get(url string, secure bool) (*http.Response, error)
	Post(url string, secure bool) (*http.Response, error)
}
type httpClient struct {
}

func NewHttpClient() *httpClient {
	return &httpClient{}
}

func (h *httpClient) Get(url string, secure bool) (*http.Response, error) {
	if !secure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *httpClient) Post(url string, secure bool, request bytes.Buffer) (*http.Response, error) {
	if !secure {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		resp, err := http.Post(url, "application/json", bytes.NewReader(request.Bytes()))
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(request.Bytes()))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
