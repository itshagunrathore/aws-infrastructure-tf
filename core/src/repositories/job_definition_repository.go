package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type JobDefinitionRepository interface {
	FindById(id uint) entities.JobDefinition
	FindByAccountIdAndJobName(accountId string, jobName string) entities.JobDefinition
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

func (repo *jobDefinitionRepository) FindByAccountIdAndJobName(accountId string, jobName string) entities.JobDefinition {
	var jobDefinition entities.JobDefinition

	repo.DB.DB().
		repo.DB.DB().Find(&jobDefinition, accountId, jobName)

	return jobDefinition
}
