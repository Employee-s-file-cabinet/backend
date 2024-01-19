package model

import "time"

type Passport struct {
	ID          uint64
	Citizenship string
	IssuedBy    string
	IssuedDate  time.Time
	Number      string
	Type        PassportType
	HasScan     bool
}

type PassportType string

const (
	PassportTypeNational      PassportType = "national"
	PassportTypeInternational PassportType = "international"
)
