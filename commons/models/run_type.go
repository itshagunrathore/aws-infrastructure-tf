package models

type RunType string

const (
	FULL       RunType = "FULL"
	DELTA      RunType = "DELTA"
	CUMULATIVE RunType = "CUMULATIVE"
)
