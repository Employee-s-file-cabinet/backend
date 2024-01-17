package model

import (
	"time"
)

type Insurance struct {
	Number  string
	HasScan bool
}

type User struct {
	ID                  uint64
	LastName            string
	FirstName           string
	MiddleName          string
	Gender              Gender
	DateOfBirth         time.Time
	PlaceOfBirth        string
	Grade               string
	PhoneNumbers        map[string]string
	Email               string
	RegistrationAddress string
	ResidentialAddress  string
	Nationality         string
	Insurance           Insurance
	TaxpayerNumber      string
	Position            string
	Department          string
}

// Gender represents user gender.
type Gender string

const (
	GenderFemale Gender = "female"
	GenderMale   Gender = "male"
)

// ExpandedUser represents summary information about the user.
type ExpandedUser struct {
	User
	Educations []Education
	Trainings  []Training
	Passports  []PassportWithVisas
	// Contracts []Contract
	// Vacations  []Vacation
}

type PassportWithVisas struct {
	Passport
	Visas []Visa
}
