package repositories

import (
	"errors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
	"gorm.io/gorm"
)

type JobDefinitionRepository interface {
	FindByAccountIdAndJobId(accountId string, jobId int) (entities.JobDefinition, error)
	FindByAccountIdAndJobName(accountId string, jobName string) (*entities.JobDefinition, error)
	Save(job entities.JobDefinition) (int, error)
}

type jobDefinitionRepository struct {
	DB db.PostgresDB
}

func NewJobDefinitionRepository(DB db.PostgresDB) JobDefinitionRepository {
	return &jobDefinitionRepository{
		DB: DB,
	}
}

func (repo *jobDefinitionRepository) FindByAccountIdAndJobId(accountId string, jobId int) (entities.JobDefinition, error) {
	var jobDefinition entities.JobDefinition
	db := repo.DB.DB()
	err := db.Where(&entities.JobDefinition{JobId: jobId, IsActive: true, IsDeleted: false}).Joins("CustomerSite", db.Where(&entities.CustomerSite{SiteId: accountId})).First(&jobDefinition).Error
	if err != nil {
		return entities.JobDefinition{}, err
	}
	return jobDefinition, nil
}

func (repo *jobDefinitionRepository) Save(job entities.JobDefinition) (int, error) {
	//TODO implement me
	result := repo.DB.DB().Save(&job)

	if result.Error != nil {
		return 0, result.Error
	}
	return job.JobId, nil
}

func (repo *jobDefinitionRepository) FindByAccountIdAndJobName(accountId string, jobName string) (*entities.JobDefinition, error) {
	var jobDefinition entities.JobDefinition
	db := repo.DB.DB()
	err := db.Where(&entities.JobDefinition{Name: jobName, IsActive: true, IsDeleted: false}).Joins("CustomerSite", db.Where(&entities.CustomerSite{SiteId: accountId})).First(&jobDefinition).Error

	// need to return some predefined error saying no record found
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &jobDefinition, nil
}
