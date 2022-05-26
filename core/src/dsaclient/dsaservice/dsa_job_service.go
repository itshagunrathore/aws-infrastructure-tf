package dsaservice

import (
	"errors"
	"fmt"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

type DsaService interface {
	CreateDsaJob(request dto.CreateDsaJobRequest) error
}

type dsaService struct {
}

func NewDsaService() DsaService {
	return &dsaService{}
}

func (d *dsaService) CreateDsaJob(request dto.CreateDsaJobRequest) error {

	host := "tdicam2118dev00.gateway.dev.cloud.teradata.com"
	dsaDriver := dsa.NewDsaDriver(host, "443", "", "", "", "")

	systemName, err := d.GetSystemName(dsaDriver)
	targetGroup, err := d.GetTarGetGroup(dsaDriver, request.SiteTargetType)
	// find a way to get this secrets
	userName := "baradmin"
	password := "V#cbjioln0367192"

	dsaRestJobPayload := models.RestJobPayload{}
	dsaRestJobPayload.RestJobDefinitionModel.JobName = request.JobName
	dsaRestJobPayload.RestJobDefinitionModel.JobType = request.JobType
	dsaRestJobPayload.RestJobDefinitionModel.JobDescription = request.Description
	dsaRestJobPayload.RestJobDefinitionModel.DataDictionaryType = "DATA"
	dsaRestJobPayload.RestJobDefinitionModel.SourceSystem = systemName
	dsaRestJobPayload.RestJobDefinitionModel.TargetGroupName = targetGroup
	dsaRestJobPayload.RestJobDefinitionModel.SrcUserName = userName
	dsaRestJobPayload.RestJobDefinitionModel.SrcUserPassword = password
	dsaRestJobPayload.RestJobSettingsModel = request.JobSettings
	dsaRestJobPayload.RestJobObjectsModels = request.JobObjects

	fmt.Println("creating dsa job")
	err = dsaDriver.PostJob(dsaRestJobPayload)

	if err != nil {
		return err
	}

	fmt.Println(systemName, targetGroup)
	// get system name
	// get target-group-name
	// get password
	// preparepaylaod
	// trigger job_creation
	// update status
	return err
}

func (d *dsaService) GetSystemName(dsaDriver dsa.DsaDriver) (string, error) {
	systems, err := dsaDriver.SystemNames()

	if err != nil {
		return "", err
	}
	var enabledSystems []models.Systems
	if systems.Valid && len(systems.Systems) > 0 {
		for _, system := range systems.Systems {
			if system.IsEnabled {
				enabledSystems = append(enabledSystems, system)
			}
		}
	}
	if len(enabledSystems) > 0 {
		return enabledSystems[0].SystemName, nil
	} else {
		return " ", errors.New("no enabled systems found")
	}
}

func (d *dsaService) GetTarGetGroup(dsaDriver dsa.DsaDriver, siteTargetType models.SiteTargetType) (string, error) {

	targetGroups, err := dsaDriver.GetTargetGroup(siteTargetType)

	if err != nil {
		return "", err
	}

	if siteTargetType == models.AWS {
		return targetGroups.S3Target[0].TargetGroupName, nil
	} else if siteTargetType == models.AZURE {
		return targetGroups.Azure[0].TargetGroupName, nil
	}
	// correct error handling needed
	return "", err
}
