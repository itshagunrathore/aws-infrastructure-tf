package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

func DsmainRestart() string {

	response := utils.DsmainRestart()
	if !response {
		fmt.Printf("DSMAIN restart failed")
	} else {
		return "DSMAIN restart Success"
	}

}
