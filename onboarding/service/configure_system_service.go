package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

// var StatusGetSystem models.DsaResponse

func ConfigureSystem(event models.Event, StatusConfigSystem *models.DetailedStatus) (string, error) {
	StatusConfigSystem.Step = "ConfigureSystem"
	var payload models.DsaSystem
	barUserSecretName := fmt.Sprintf("%s_TDaaS_BAR", event.accountId)
	barUserPassword, err := utils.GetSecret(barUserSecretName, event.Region, event.cloudPlatform)
	if err != nil {
		log.Error(err)
		return "Failed to fetch baruser password", err
	}
	payload.Password = barUserPassword
	payload.User = "TDaaS_BAR"
	payload.SkipForceFull = true
	// payload.SoftLimit = 1
	payload.SystemName = event.SystemName
	payload.TdpID = event.SystemName
	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata", event.DscIp, event.Port)
	response, err := dsa.PostConfigDsc(url, payload, &StatusConfigSystem)
	if err != nil {
		StatusConfigSystem.Error = err
		StatusConfigSystem.StepStatus = "Failed"
		// StatusConfigSystem.StatusCode = int(response[])
		return "Failed to configure system on dsc", err
	} else {
		var concfigsystemresponse models.ConfigSystemResponse
		json.Unmarshal(response, &concfigsystemresponse)
		StatusConfigSystem.StepResponse = concfigsystemresponse.Status
		StatusConfigSystem.StepStatus = "Success"
		return string(response), err
	}
}

func GetSystemName(event event, StatusGetSystem *models.DetailedStatus) ([]string, error) {
	StatusGetSystem.subStep = "GetSystemName"
	url := fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata", event.DscIp, event.Port)
	response, err := dsa.GetConfigDsc(url, &StatusGetSystem)
	if err != nil {
		fmt.Printf("\nResponse:%v\n", StatusGetSystem)
		StatusGetSystem.StepStatus = "Failed"
		StatusGetSystem.StepResponse = string(response)
		StatusGetSystem.Error = err
		StatusGetSystem.StatusCode = 500
		return []string{"Failed to fetch media servers", ""}, err
	}
	var systems models.GetSystemNames
	var Pog string
	json.Unmarshal(response, &systems)
	for _, system := range systems.Systems {
		Pog = system.SystemName
	}

	url = fmt.Sprintf("https://%s:%s/dsa/components/systems/teradata/%s", event.DscIp, event.Port, Pog)
	response, err = dsa.GetConfigDsc(url, &StatusGetSystem)
	var getSystemResponse models.GetSystem
	json.Unmarshal(response, &getSystemResponse)
	data, _ := json.Marshal(getSystemResponse)
	fmt.Println(string(data))
	var PogIps []string
	for _, ip := range getSystemResponse.Nodes {
		PogIps = append(PogIps, ip.IPAddress[1])
	}

	return PogIps, err
}
