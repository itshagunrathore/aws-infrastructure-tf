package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func UpdateMediaServers(event models.Event, StatusUpdateMedia *models.DetailedStatus) ([]string, error) {
	url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
	response, err := dsa.GetConfigDsc(url, &StatusUpdateMedia)
	if err != nil {
		StatusUpdateMedia.StepResponse = "Failed to fetch media servers"
		return []string{"Failed to fetch media servers", ""}, err
	}
	var mediaPayload models.MediaServersConfig
	var mediaResponse models.MediaResponse
	var ipinfo models.IPInfo
	json.Unmarshal(response, &mediaResponse)
	data, _ := json.Marshal(mediaResponse)
	log.Info(string(data))
	for _, media := range mediaResponse.Medias {
		mediaPayload.ServerName = media.ServerName
		mediaPayload.PoolSharedPipes = 100
		mediaPayload.Port = 15401
		for _, netmask := range media.Ips {
			ipinfo.IPAddress = netmask.IPAddress
			ipinfo.Netmask = "255.255.255.255"
			fmt.Printf("\nCurrent Media ip:%v\n", ipinfo)
			mediaPayload.IPInfo = append(mediaPayload.IPInfo, ipinfo)
			fmt.Printf("\nCurrent Media payload ip:%v\n", mediaPayload.IPInfo)
			ipinfo = models.IPInfo{}
		}
		url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
		var configmediaresponse models.ConfigMediaResponse
		response, err := dsa.PostConfigDsc(url, mediaPayload, &StatusUpdateMedia)
		if err != nil {
			StatusUpdateMedia.StepStatus = "Failed"
			StatusUpdateMedia.StepResponse = mediaResponse.Status
			StatusUpdateMedia.Error = err
			StatusUpdateMedia.StatusCode = 500
			return []string{"Failed to configure target group"}, err
		} else {
			json.Unmarshal(response, &configmediaresponse)
			log.Info(response)
			StatusUpdateMedia.StepStatus = "Success"
			StatusUpdateMedia.StepResponse = configmediaresponse.Status
			StatusUpdateMedia.Error = err
			// StatusUpdateMedia.StatusCode = 200
		}
		mediaPayload = models.MediaServersConfig{}
	}
	log.Info("Media server response:%v\n", mediaResponse)
	return []string{"Media server updated"}, err
}

func GetMedia(event models.Event, StatusGetMedia *models.DetailedStatus) ([]string, error) {
	PogIps, err := dsa.GetSystemName(event, &StatusGetMedia)

	url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
	response, err := dsa.GetConfigDsc(url, &StatusGetMedia)
	if err != nil {
		log.Info("Failed to fetch media servers")
		return []string{"Failed to fetch media servers", ""}, err
	}
	var mediaResponse models.MediaResponse

	json.Unmarshal(response, &mediaResponse)
	data, _ := json.Marshal(mediaResponse)
	log.Info("Fetched media response: " + string(data))

	var LiveMediaServer []string
	for _, media := range mediaResponse.Medias {
		for _, v := range media.Ips {
			for j := 0; j < len(PogIps); j++ {
				if PogIps[j] == v.IPAddress {
					LiveMediaServer = append(LiveMediaServer, media.ServerName)
				}

			}
		}

	}
	return LiveMediaServer, err
}
