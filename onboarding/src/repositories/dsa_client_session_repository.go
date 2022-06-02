package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
)

type DsaClientSessionRepository interface {
	Get() (entities.DsaClientSession, error)
	Post(entities.DsaClientSession) error
	Update(entities.DsaClientSession) error
}
type dsaClientSessionRepository struct {
	DB db.PostgresDB
}

func NewDsaClientSessionRepository(DB db.PostgresDB) DsaClientSessionRepository {
	return &dsaClientSessionRepository{
		DB: DB,
	}
}

//this will return the latest row in the database
func (d *dsaClientSessionRepository) Get() (entities.DsaClientSession, error) {
	var dsaClientSessionEntity entities.DsaClientSession
	err := d.DB.DB().Table(dsaClientSessionEntity.TableName()).Last(&dsaClientSessionEntity).Error
	if err != nil {
		return entities.DsaClientSession{}, err
	}
	return dsaClientSessionEntity, nil
}

func (d *dsaClientSessionRepository) Post(e entities.DsaClientSession) error {
	err := d.DB.DB().Create(&e).Error
	// check if we need to commit
	if err != nil {
		return err
	}
	return nil
}
func (d *dsaClientSessionRepository) Update(e entities.DsaClientSession) error {
	err := d.DB.DB().Save(&e).Error
	// check if we need to commit
	if err != nil {
		return err
	}
	return nil
}
