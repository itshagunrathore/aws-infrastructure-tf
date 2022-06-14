package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
)

type LatestJobSessionRepository interface {
	Save(latestJobSession entities.LatestJobSession) (int, error)
}

type latestJobSessionRepository struct {
	DB db.PostgresDB
}

func NewLatestJobSessionRepository(DB db.PostgresDB) LatestJobSessionRepository {
	return &latestJobSessionRepository{
		DB: DB,
	}
}

func (l latestJobSessionRepository) Save(latestJobSession entities.LatestJobSession) (int, error) {
	//TODO implement me
	panic("implement me")
}
