package repositories

import (
	"errors"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
	"gorm.io/gorm"
)

type JobDefinitionRepository interface {
	FindById(id uint) entities.JobDefinition
	FindByAccountIdAndJobName(accountId string, jobName string) (*entities.JobDefinition, error)
}

type jobDefinitionRepository struct {
	DB db.PostgresDB
}

func NewJobDefinitionRepository(DB db.PostgresDB) JobDefinitionRepository {
	return &jobDefinitionRepository{
		DB: DB,
	}
}

func (repo *jobDefinitionRepository) FindById(id uint) entities.JobDefinition {
	var jobDefinition entities.JobDefinition
	repo.DB.DB().Find(&jobDefinition, id)

	return jobDefinition
}

func (repo *jobDefinitionRepository) FindByAccountIdAndJobName(accountId string, jobName string) (*entities.JobDefinition, error) {
	var jobDefinition entities.JobDefinition
	db := repo.DB.DB()
	err := db.Joins("CustomerSite", db.Where(&entities.CustomerSite{SiteId: accountId})).First(&jobDefinition, "jobDefinition.Name = ?",
		jobName, "jobDefinition.IsDeleted = ?", false, "").Error

	// need to return some predefined error saying no record found
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &jobDefinition, nil
}
