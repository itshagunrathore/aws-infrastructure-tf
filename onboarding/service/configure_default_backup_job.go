package service

import (
	"encoding/json"
	"fmt"

	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
)

func CreateDefaulJob(event models.Event, StatusAWSApp *models.DetailedStatus) (string, error) {
	StatusAWSApp.Step = "CreateDefaultJob"
	var payload models.CreateJob
	var backupObjects models.RestJobObjectsModels
	var createjobresponse models.CreateJobResponse
	backupObjects.IncludeAll = false
	backupObjects.ObjectName = "DBC"
	backupObjects.ObjectType = "DATABASE"
	backupObjects.ParentType = "BACKUP_JOB"
	payload.RestJobDefinitionModel.DataDictionaryType = "DATA"
	payload.RestJobDefinitionModel.JobName = fmt.Sprintf("%s_default_dbc_only", event.SystemName)
	payload.RestJobDefinitionModel.JobType = "BACKUP"
	payload.RestJobDefinitionModel.SourceSystem = event.SystemName
	payload.RestJobDefinitionModel.SrcUserName = event.Db_user
	payload.RestJobDefinitionModel.SrcUserPassword = event.Db_password
	payload.RestJobDefinitionModel.TargetGroupName = "TG_BAAS_GO3"
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
	response, err := dsa.PostConfigDsc(url, payload, &StatusAWSApp)
	if err != nil {
		data := json.Unmarshal(response, &createjobresponse)
		fmt.Printf("DSA output:%v", data)
		StatusAWSApp.StepStatus = "Failed"
		StatusAWSApp.StepResponse = createjobresponse.Status
		return "Failed to configure target group", err
	} else {
		StatusAWSApp.StepStatus = "Success"
		return string(response), err
	}

}
