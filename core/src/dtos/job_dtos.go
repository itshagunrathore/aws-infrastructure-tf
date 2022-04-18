package dtos

// some values are missing
type GetAllJobsDto struct {
	Job_Id                 int    `json:"job_id"`
	Name                   string `json:"name" gorm:"column:job_name"`
	Description            string `json:"description"`
	Site_Id                string `json:"site_id" gorm:"column:site_id"`
	Site_Target_Type       string `json:"site_target_type"`
	Is_Active              bool   `json:"is_active"`
	Job_Priority           int    `json:"job_priority"`
	Job_Type               string `json:"job_type"` // need to check on enums
	Backup_Type            string `json:"backup_type"`
	No_Of_Retention_Copies int    `json:"no_of_retention_copies" gorm:"column:retention_copies_count"`
	Is_Auto_Abort_Active   bool   `json:"is_auto_abort_active"`
	Auto_Abort_In_Min      int    `json:"auto_abort_in_min"`
}
