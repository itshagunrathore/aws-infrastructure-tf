package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type JobDefinitionRepository interface {
	FindById(id uint) entities.JobDefinition
}

type jobDefinitionRepository struct {
	DB db.PostgresDB
}

func NewJobDefinitionRepository(DB db.PostgresDB) JobDefinitionRepository {
	return &jobDefinitionRepository{
		DB: DB,
	}
}

func (j *jobDefinitionRepository) FindById(id uint) entities.JobDefinition {
	var jobDefinition entities.JobDefinition
	j.DB.DB().Find(&jobDefinition, id)

	return jobDefinition
}
