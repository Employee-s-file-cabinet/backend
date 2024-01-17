package postgresql

import (
	"time"

	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
)

type shortUserInfo struct {
	ID                uint64 `db:"id"`
	Department        string `db:"department"`
	Email             string `db:"work_email"`
	FirstName         string `db:"firstname"`
	LastName          string `db:"lastname"`
	MiddleName        string `db:"middlename"`
	MobilePhoneNumber string `db:"mobile_phone_number"`
	OfficePhoneNumber string `db:"office_phone_number"`
	Position          string `db:"position"`
}

type user struct {
	shortUserInfo
	Gender              gender    `db:"gender"`
	DateOfBirth         time.Time `db:"date_of_birth"`
	PlaceOfBirth        string    `db:"place_of_birth"`
	Grade               string    `db:"grade"`
	RegistrationAddress string    `db:"registration_address"`
	ResidentialAddress  string    `db:"residential_address"`
	Nationality         string    `db:"nationality"`
	InsuranceNumber     string    `db:"insurance_number"`
	TaxpayerNumber      string    `db:"taxpayer_number"`
	PositionID          uint64    `db:"position_id"`
	DepartmentID        uint64    `db:"department_id"`
}

type gender string

const (
	genderFemale gender = "Женский"
	genderMale   gender = "Мужской"
)

func convertShortUserInfoToModelShortUserInfo(info shortUserInfo) model.ShortUserInfo {
	return model.ShortUserInfo(info)
}

func convertUserToModelUser(user *user) model.User {
	mu := model.User{
		ShortUserInfo:       convertShortUserInfoToModelShortUserInfo(user.shortUserInfo),
		DateOfBirth:         user.DateOfBirth,
		PlaceOfBirth:        user.PlaceOfBirth,
		Grade:               user.Grade,
		RegistrationAddress: user.RegistrationAddress,
		ResidentialAddress:  user.ResidentialAddress,
		Nationality:         user.Nationality,
		InsuranceNumber:     user.InsuranceNumber,
		TaxpayerNumber:      user.TaxpayerNumber,
		PositionID:          user.PositionID,
		DepartmentID:        user.DepartmentID,
	}
	switch user.Gender {
	case genderMale:
		mu.Gender = model.GenderMale
	case genderFemale:
		mu.Gender = model.GenderFemale
	}
	return mu
}

func convertModelUserToUser(u *model.User) user {
	var gr gender
	switch u.Gender {
	case model.GenderMale:
		gr = genderMale
	case model.GenderFemale:
		gr = genderFemale
	}

	return user{
		shortUserInfo:       shortUserInfo(u.ShortUserInfo),
		Gender:              gr,
		DateOfBirth:         u.DateOfBirth,
		PlaceOfBirth:        u.PlaceOfBirth,
		Grade:               u.Grade,
		RegistrationAddress: u.RegistrationAddress,
		ResidentialAddress:  u.ResidentialAddress,
		Nationality:         u.Nationality,
		InsuranceNumber:     u.InsuranceNumber,
		TaxpayerNumber:      u.TaxpayerNumber,
		PositionID:          u.PositionID,
		DepartmentID:        u.DepartmentID,
	}
}

type listUser struct {
	shortUserInfo
	TotalCount int `db:"total_count"`
}

type education struct {
	ID                uint64    `db:"id"`
	Number            string    `db:"document_number"`
	Program           string    `db:"title_of_program"`
	IssuedInstitution string    `db:"title_of_institution"`
	DateTo            time.Time `db:"year_of_end"`
	DateFrom          time.Time `db:"year_of_begin"`
}

func convertEducationToModelEducation(ed education) model.Education {
	return model.Education(ed)
}

type training struct {
	ID                uint64    `db:"id"`
	Program           string    `db:"title_of_program"`
	IssuedInstitution string    `db:"title_of_institution"`
	Cost              uint64    `db:"cost"`
	DateTo            time.Time `db:"date_end"`
	DateFrom          time.Time `db:"date_begin"`
}

func convertTrainingToModelTraining(tr training) model.Training {
	return model.Training(tr)
}

type passport struct {
	ID         uint64       `db:"id"`
	IssuedBy   string       `db:"issued_by"`
	IssuedDate time.Time    `db:"issued_date"`
	Number     string       `db:"number"`
	Type       passportType `db:"type"`
}

type passportType string

const (
	passportTypeExternal   passportType = "Заграничный"
	passportTypeForeigners passportType = "Иностранного гражданина"
	passportTypeInternal   passportType = "Внутренний"
)

func convertPassportToModelPassport(p passport) model.Passport {
	var pt model.PassportType
	switch p.Type {
	case passportTypeExternal:
		pt = model.PassportTypeExternal
	case passportTypeInternal:
		pt = model.PassportTypeInternal
	case passportTypeForeigners:
		pt = model.PassportTypeForeigners
	}

	return model.Passport{
		ID:         p.ID,
		IssuedBy:   p.IssuedBy,
		IssuedDate: p.IssuedDate,
		Number:     p.Number,
		Type:       pt,
	}
}

func convertModelPassportToPassport(mp model.Passport) passport {
	var t passportType
	switch mp.Type {
	case model.PassportTypeExternal:
		t = passportTypeExternal
	case model.PassportTypeInternal:
		t = passportTypeInternal
	case model.PassportTypeForeigners:
		t = passportTypeForeigners
	}

	return passport{
		ID:         mp.ID,
		IssuedBy:   mp.IssuedBy,
		IssuedDate: mp.IssuedDate,
		Number:     mp.Number,
		Type:       t,
	}
}

type visa struct {
	ID            uint64                  `db:"id"`
	Number        string                  `db:"number"`
	Type          visaType                `db:"type"`
	IssuedState   string                  `db:"issued_state"`
	ValidTo       time.Time               `db:"valid_to"`
	ValidFrom     time.Time               `db:"valid_from"`
	NumberEntries model.VisaNumberEntries `db:"number_entries"`
}

type visaType string

const (
	visaTypeExternal   passportType = "Заграничная"
	visaTypeForeigners passportType = "Иностранного гражданина"
)

func convertVisaToModelVisa(v visa) model.Visa {
	return model.Visa{
		ID:            v.ID,
		Number:        v.Number,
		IssuedState:   v.IssuedState,
		ValidTo:       v.ValidTo,
		ValidFrom:     v.ValidFrom,
		NumberEntries: v.NumberEntries,
	}
}

func convertModelVisaToVisa(mv model.Visa) visa {
	return visa{
		ID:            mv.ID,
		Number:        mv.Number,
		IssuedState:   mv.IssuedState,
		ValidTo:       mv.ValidTo,
		ValidFrom:     mv.ValidFrom,
		NumberEntries: mv.NumberEntries,
	}
}

type vacation struct {
	ID        uint64    `db:"id"`
	DateBegin time.Time `db:"date_begin"`
	DateEnd   time.Time `db:"date_end"`
}

func convertVacationToModelVacation(v vacation) model.Vacation {
	return model.Vacation(v)
}
