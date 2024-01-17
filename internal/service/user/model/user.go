package model

import (
	"time"
)

type ShortUserInfo struct {
	ID                uint64
	Department        string
	Email             string
	FirstName         string
	LastName          string
	MiddleName        string
	MobilePhoneNumber string
	OfficePhoneNumber string
	Position          string
}

type User struct {
	ShortUserInfo
	Gender              gender
	DateOfBirth         time.Time
	PlaceOfBirth        string
	Grade               string
	RegistrationAddress string
	ResidentialAddress  string
	Nationality         string
	InsuranceNumber     string
	TaxpayerNumber      string
	PositionID          uint64
	DepartmentID        uint64
}

// gender represents user gender.
type gender string

const (
	GenderFemale gender = "female"
	GenderMale   gender = "male"
)

// ExpandedUser represents summary information about the user.
type ExpandedUser struct {
	User
	Educations        []Education
	Trainings         []Training
	Passports         []Passport
	ExternalPassports []Passport
	Visas             []Visa
	ForeignersVisas   []Visa
	// Contracts []Contract
	Vacations []Vacation
}
