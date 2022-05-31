package dsa

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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

func PostConfigDsc(url string, payload interface{}) (string, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	data, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		// msg := fmt.Sprintf("Unable to make PUT request for pod account service to set the active and inactive system for upgrading account %s: %v", accountId, err)
		return "Failed", err
	}
	req.Header.Set("Content-Type", "application/json")
	msg := "Hey there"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "failed", errors.New(msg)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)

		return "failed", errors.New(msg)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(response), err
}

func GetConfigDsc(url string) ([]byte, error) {
	fmt.Println(url)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Failed to create http request")
		msg := "Failed to create http request"
		return []byte(msg), err
	}

	msg := "Hey there"
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Failed to invoke dsa %s", err)

	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		fmt.Printf("Issue in invocation %s", err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200 {
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Printf("Http non 200 response %s", err)

		return response, errors.New(msg)
	}
	return response, err
}
