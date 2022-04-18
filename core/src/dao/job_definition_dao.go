package dao

import (
	"baas/src/database"
	"baas/src/dtos"
)

var DB = database.GetDatabase()

func GetAllJobs(site_id string) (results []dtos.GetAllJobsDto) {
	newResults := []dtos.GetAllJobsDto{}
	_ = DB.Table("job_definition").Select("*").Joins("JOIN customer_site on customer_site.customer_site_id = job_definition.customer_site_id").
		Find(&newResults, "customer_site.site_id = ?", site_id)

	return newResults
}
