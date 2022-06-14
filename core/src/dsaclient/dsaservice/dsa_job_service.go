package dsaservice

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/drivers/dsa"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/models"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/dto"
)

type DsaJobService interface {
	CreateDsaJob(context *gin.Context, request dto.CreateDsaJobRequest) error
}

type dsaJobService struct {
}

func NewDsaService() DsaJobService {
	return &dsaJobService{}
}

func (d *dsaJobService) CreateDsaJob(context *gin.Context, request dto.CreateDsaJobRequest) error {
	// TODO START DSA HERE
	host := "tdicam2118dev00.gateway.dev.cloud.teradata.com"
	dsaDriver := dsa.NewDsaDriver(host, "443", "", "", "", "")

	systemName, err := d.GetSystemName(dsaDriver)
	if err != nil {
		log.Errorw(err.Error())
		return err
		// TODO log this error and publish msg on kafka regarding failure are return from this function and find a way to stop go routine
	}
	targetGroup, err := d.GetTarGetGroup(dsaDriver, request.SiteTargetType)
	if err != nil {
		return err
	}
	// TODO for antares fetch secrets from secrets manager for time being using baradmin user
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
	dsaRestJobPayload.RestJobSettingsModel.Online = true
	dsaRestJobPayload.RestJobObjectsModels = request.JobObjects
	err = dsaDriver.PostJob(dsaRestJobPayload)
	// Todo to push the status to dsa-notifications kafka topic update status accordingly in baas-db in that function create entry in latest job session
	// status to be success
	// kafka notification
	// Todo to stop dsa
	if err != nil {
		return err
	}
	fmt.Println(systemName, targetGroup)

	return err
}

func (d *dsaJobService) GetSystemName(dsaDriver dsa.DsaDriver) (string, error) {
	systems, err := dsaDriver.SystemNames()
	log.Infow("systemname call recived")
	if err != nil {
		return "", errors.New("no enabled system found")
	}
	var enabledSystems []models.Systems
	if systems.Valid && len(systems.Systems) > 0 {
		for _, system := range systems.Systems {
			if system.IsEnabled {
				enabledSystems = append(enabledSystems, system)
			}
		}
	}
	fmt.Println("=======================================")
	fmt.Println(enabledSystems)
	if len(enabledSystems) > 0 {
		return enabledSystems[0].SystemName, nil
	}
	return "", errors.New("no enabled system found")
}

func (d *dsaJobService) GetTarGetGroup(dsaDriver dsa.DsaDriver, siteTargetType models.SiteTargetType) (string, error) {

	targetGroups, err := dsaDriver.GetTargetGroup(siteTargetType)
	fmt.Println("=======================================")
	fmt.Println(targetGroups)
	if err != nil {
		return "", errors.New("error while fetching target groups")
	}
	// TODO currently we are going with vantage assumption once we decide on fix target-group-name for baas we will check for that
	if siteTargetType == models.AWS {
		return targetGroups.S3Target[0].TargetGroupName, nil
	} else if siteTargetType == models.AZURE {
		return targetGroups.Azure[0].TargetGroupName, nil
	} else if siteTargetType == models.GCP {
		return targetGroups.Gcp[0].TargetGroupsName, nil
	}

	return "", errors.New("no target-group found for job creation")
}
