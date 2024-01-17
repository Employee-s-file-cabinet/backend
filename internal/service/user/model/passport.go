package model

import "time"

type Passport struct {
	ID           uint64
	Number       string
	Type         passportType
	Citizenship  string
	IssuedBy     string
	IssuedByCode string
	IssuedDate   time.Time
	ExpiredAt    time.Time
}

type passportType string

const (
	PassportTypeExternal passportType = "external"
	PassportTypeInternal passportType = "internal"
)
