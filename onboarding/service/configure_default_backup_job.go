package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/utils"
)

func CreateDefaulJob(event models.Event, StatusAWSApp *models.DetailedStatus) (string, error) {
	StatusAWSApp.Step = "CreateDefaultJob"
	var payload models.CreateJob
	var backupObjects models.RestJobObjectsModels
	var createJobResponse models.CreateJobResponse
	barUserSecretName := fmt.Sprintf("%s_TDaaS_BAR", event.AccountId)
	BarUserPassword, err := utils.GetSecret(barUserSecretName, event.Region, event.CloudPlatform)
	backupObjects.IncludeAll = false
	backupObjects.ObjectName = "DBC"
	backupObjects.ObjectType = "DATABASE"
	backupObjects.ParentType = "BACKUP_JOB"
	payload.RestJobDefinitionModel.DataDictionaryType = "DATA"
	payload.RestJobDefinitionModel.JobName = fmt.Sprintf("%s_default_dbc_only", event.SystemName)
	payload.RestJobDefinitionModel.JobType = "BACKUP"
	payload.RestJobDefinitionModel.SourceSystem = event.SystemName
	payload.RestJobDefinitionModel.SrcUserName = "TDaaS_BAR"
	payload.RestJobDefinitionModel.SrcUserPassword = BarUserPassword
	payload.RestJobDefinitionModel.TargetGroupName = "TG_BAAS"
	payload.RestJobObjectsModels = append(payload.RestJobObjectsModels, backupObjects)
	payload.RestJobSettingsModel.BlockLevelCompression = "DEFAULT"
	payload.RestJobSettingsModel.LoggingLevel = "Error"
	payload.RestJobSettingsModel.Nosync = false
	payload.RestJobSettingsModel.Nowait = false
	payload.RestJobSettingsModel.Online = false
	payload.RestJobSettingsModel.Reblock = false
	payload.RestJobSettingsModel.SkipArchive = false
	payload.RestJobSettingsModel.SkipJoinhashIndex = false
	payload.RestJobSettingsModel.SkipStats = false
	payload.RestJobSettingsModel.TemperatureOverride = "DEFAULT"
	payload.RestJobSettingsModel.TrackEmptyTables = false

	log.Info(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/jobs", event.DscIp, event.Port)
	log.Info("invoking dsa api: %s", url)
	response, err := dsa.PostConfigDsc(url, payload, &StatusAWSApp)
	if err != nil {
		data := json.Unmarshal(response, &createJobResponse)
		log.Info("DSA response:%v", data)
		StatusAWSApp.StepStatus = "Failed"
		StatusAWSApp.StepResponse = createJobResponse.Status
		return "Failed to configure target group", err
	} else {
		StatusAWSApp.StepStatus = "Success"
		return string(response), err
	}

}
