package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/onboarding/entities"
)

type CustomerSiteRepository interface {
	Get(accountId string) (entities.CustomerSite, error)
}

type customerSiteRepository struct {
	DB db.PostgresDB
}

func NewCustomerSiteRepository(DB db.PostgresDB) CustomerSiteRepository {
	return &customerSiteRepository{
		DB: DB,
	}
}

func (c customerSiteRepository) Get(accountId string) (entities.CustomerSite, error) {
	var customerSiteEntity entities.CustomerSite
	err := c.DB.DB().Where("site_id = ?", accountId).First(&customerSiteEntity).Error

	if err != nil {
		return entities.CustomerSite{}, err
	}
	return customerSiteEntity, nil
}

func (repo customerSiteRepository) OnboardNewTenant(NewTenant entities.CustomerSite) {
	repo.DB.create(&NewTenant)

}
