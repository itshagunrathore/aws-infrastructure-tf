package dsa

import (
	"bytes"
	"crypto/tls"
	"net/http"
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
