package entities

import "time"

//assuming that we are using snake case for db colums
type DsaClientSession struct {
	ClientSessionId string    `gorm:"type:varchar;column:client_session_id;not null"`
	TimeCreated     time.Time `gorm:"column:time_created;not null"`
	TimeUpdated     time.Time `gorm:"column:time_updated;not null"`
	IsDeleted       bool      `gorm:"is_deleted;default=false"`
	AccountId       string    `gorm:"account_id;not null"`
}

func (d *DsaClientSession) TableName() string {
	return "dsa_client_session"
}
