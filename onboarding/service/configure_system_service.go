package service

import (
	"encoding/json"
	"fmt"
)

func ConfigureSystem(event event) (string, error) {
	var payload DsaSystem
	payload.Password = event.Db_password
	payload.User = event.Db_user
	payload.SkipForceFull = true
	// payload.SoftLimit = 1
	payload.SystemName = event.SystemName
	payload.TdpID = event.SystemName
	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata", event.DscIp, event.Port)
	response, err := PostConfigDsc(url, payload)
	if err != nil {
		return "Failed to configure target group", err
	} else {
		return string(response), err
	}
}

func GetSystemName(event event) ([]string, error) {
	url := fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata", event.DscIp, event.Port)
	response, err := GetConfigDsc(url, event.DscIp, event.Port)

	var systems GetSystemNames
	var Pog string
	json.Unmarshal(response, &systems)
	for _, system := range systems.Systems {
		Pog = system.SystemName
	}

	url = fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata/%s", event.DscIp, event.Port, Pog)
	response, err = GetConfigDsc(url, event.DscIp, event.Port)
	var jresponse GetSystem
	json.Unmarshal(response, &jresponse)
	var PogIps []string
	for _, ip := range jresponse.Nodes {
		PogIps = append(PogIps, ip.IPAddress[1])
	}

	return PogIps, err
}
