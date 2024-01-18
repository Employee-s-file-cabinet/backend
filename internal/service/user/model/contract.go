package model

import "time"

type Contract struct {
	ID              uint64
	Number          string
	HasScan         bool
	ContractType    ContractType
	WorkTypeID      uint64
	WorkType        string
	ProbationPeriod uint32
	DateBegin       time.Time
	DateEnd         time.Time
}

type ContractType string

const (
	ContractTypePermanent ContractType = "permanent"
	ContractTypeTemporary ContractType = "temporary"
)
