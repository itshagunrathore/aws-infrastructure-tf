package models

type JobType string

const (
	Restore     JobType = "RESTORE"
	Replication JobType = "REPLICATION"
	Backup      JobType = "BACKUP"
)
