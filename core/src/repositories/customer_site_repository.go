package repositories

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/db"
<<<<<<< HEAD
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/onboarding/entities"
=======
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/entities"
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
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
<<<<<<< HEAD

func (repo customerSiteRepository) OnboardNewTenant(NewTenant entities.CustomerSite) {
	repo.DB.create(&NewTenant)

}
=======
>>>>>>> cb617580f7e540b5109a595dbdc81d6aa6c40d39
