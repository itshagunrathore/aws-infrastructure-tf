package service

import "fmt"

func CreateDefaulJob(event event) (string, error) {
	var payload CreateJob
	var backupObjects RestJobObjectsModels
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

	fmt.Println(payload)
	url := fmt.Sprintf("https://%s:%s/dsa/jobs", event.DscIp, event.Port)
	response, err := PostConfigDsc(url, payload)
	if err != nil {
		return "Failed to configure target group", err
	} else {
		return string(response), err
	}

}
