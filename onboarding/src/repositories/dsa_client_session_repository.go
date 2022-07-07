package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding/src/entities"
)

type DsaClientSessionRepository interface {
	GetProvisionedAccounts(dsaClientSessionEntity *entities.DsaClientSession) (entities.DsaClientSession, error)
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
func (d *dsaClientSessionRepository) GetProvisionedAccounts(dsaClientSessionEntity *entities.DsaClientSession) (entities.DsaClientSession, error) {
	var resp entities.DsaClientSession
	err := d.DB.DB().Where("account_id = ? and is_deleted = ?", dsaClientSessionEntity.AccountId, dsaClientSessionEntity.IsDeleted).First(&resp).Error
	if err != nil {
		return resp, err
	}
	return resp, nil
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
	err := d.DB.DB().Where("client_session_id = ?", e.ClientSessionId).UpdateColumns(&e).Error
	// check if we need to commit
	if err != nil {
		return err
	}
	return nil
}