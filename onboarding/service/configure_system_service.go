package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

func ConfigureSystem(event models.Event, StatusConfigSystem *models.DetailedStatus) (string, error) {
	StatusConfigSystem.Step = "ConfigureSystem"
	var payload models.DsaSystem
	barUserSecretName := fmt.Sprintf("%s_TDaaS_BAR", event.AccountId)
	barUserPassword, err := utils.GetSecret(barUserSecretName, event.Region, event.CloudPlatform)
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
	log.Info(payload)
	url := fmt.Sprintf("https://%s:%s%s", event.DscIp, event.Port, models.ConfigSystem)
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

func GetSystemName(event models.Event, StatusGetSystem *models.DetailedStatus) ([]string, error) {
	StatusGetSystem.SubStep = "GetSystemName"
	url := fmt.Sprintf("https://%s:%s%s", event.DscIp, event.Port, models.ConfigSystem)
	response, err := dsa.GetConfigDsc(url, &StatusGetSystem)
	if err != nil {
		log.Info("\nResponse:%v\n", StatusGetSystem)
		StatusGetSystem.StepStatus = "Failed"
		StatusGetSystem.StepResponse = string(response)
		StatusGetSystem.Error = err
		StatusGetSystem.StatusCode = 500
		return []string{"Failed to fetch media servers", ""}, err
	}
	var systems models.GetSystemNames
	var Pog string
	json.Unmarshal(response, &systems)
	// User System name from payload
	for _, system := range systems.Systems {
		Pog = system.SystemName
	}

	url = fmt.Sprintf("https://%s:%s%s/%s", event.DscIp, event.Port, models.ConfigSystem, Pog)
	response, err = dsa.GetConfigDsc(url, &StatusGetSystem)
	var getSystemResponse models.GetSystem
	json.Unmarshal(response, &getSystemResponse)
	data, _ := json.Marshal(getSystemResponse)
	log.Info(string(data))
	var PogIps []string
	for _, ip := range getSystemResponse.Nodes {
		PogIps = append(PogIps, ip.IPAddress[1])
	}

	return PogIps, err
}
