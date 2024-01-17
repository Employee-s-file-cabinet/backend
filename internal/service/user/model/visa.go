package model

import "time"

type Visa struct {
	ID            uint64
	Number        string
	Type          visaType
	Category      string
	IssuedState   string
	IssuedDate    time.Time
	ExpiredAtDate time.Time
}

type visaType string

const (
	VisaTypeExternal   visaType = "external"
	VisaTypeForeigners visaType = "foreigners"
)
