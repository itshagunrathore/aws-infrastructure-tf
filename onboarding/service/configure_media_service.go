package service

import (
	"encoding/json"
	"fmt"
)

func UpdateMediaServers(event event) ([]string, error) {
	url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
	response, err := GetConfigDsc(url, event.DscIp, event.Port)
	if err != nil {
		return []string{"Failed to fetch media servers", ""}, err
	}
	var mediaPayload MediaServersConfig
	var jresponse MediaResponse
	var ipinfo IPInfo
	json.Unmarshal(response, &jresponse)
	for _, media := range jresponse.Medias {
		mediaPayload.ServerName = media.ServerName
		mediaPayload.PoolSharedPipes = 100
		mediaPayload.Port = 15401
		for _, netmask := range media.Ips {
			ipinfo.IPAddress = netmask.IPAddress
			ipinfo.Netmask = "255.255.255.255"
			fmt.Printf("\nCurrent Media ip:%v\n", ipinfo)
			mediaPayload.IPInfo = append(mediaPayload.IPInfo, ipinfo)
			fmt.Printf("\nCurrent Media payload ip:%v\n", mediaPayload.IPInfo)
			ipinfo = IPInfo{}
		}
		url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
		response, err := PostConfigDsc(url, mediaPayload)
		if err != nil {
			return []string{"Failed to configure target group"}, err
		} else {
			fmt.Println(response)
		}
		mediaPayload = MediaServersConfig{}
	}
	fmt.Printf("Value of media servers:%v\n", jresponse)
	return []string{"Media updated"}, err
}

func GetMedia(event event) ([]string, error) {
	PogIps, err := GetSystemName(event)

	url := fmt.Sprintf("https://%s:%s/dsa/components/mediaservers", event.DscIp, event.Port)
	response, err := GetConfigDsc(url, event.DscIp, event.Port)
	if err != nil {
		return []string{"Failed to fetch media servers", ""}, err
	}
	var jresponse MediaResponse

	json.Unmarshal(response, &jresponse)

	var LiveMediaServer []string
	for _, media := range jresponse.Medias {
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
