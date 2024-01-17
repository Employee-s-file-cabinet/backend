package postgresql

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
)

type user struct {
	ID                  uint64       `db:"id"`
	LastName            string       `db:"lastname"`
	FirstName           string       `db:"firstname"`
	MiddleName          string       `db:"middlename"`
	Gender              gender       `db:"gender"`
	DateOfBirth         time.Time    `db:"date_of_birth"`
	PlaceOfBirth        string       `db:"place_of_birth"`
	Grade               string       `db:"grade"`
	PhoneNumbers        phoneNumbers `db:"phone_numbers"`
	Email               string       `db:"work_email"`
	RegistrationAddress string       `db:"registration_address"`
	ResidentialAddress  string       `db:"residential_address"`
	Nationality         string       `db:"nationality"`
	InsuranceNumber     string       `db:"insurance_number"`
	InsuranceHasScan    bool         `db:"insurance_has_scan"`
	TaxpayerNumber      string       `db:"taxpayer_number"`
	Position            string       `db:"position"`
	Department          string       `db:"department"`
}

type gender string

const (
	genderFemale gender = "Женский"
	genderMale   gender = "Мужской"
)

type phoneNumbers map[string]string

func (ph *phoneNumbers) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &ph)
	case string:
		return json.Unmarshal([]byte(v), &ph)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}
func (ph *phoneNumbers) Value() (driver.Value, error) {
	return json.Marshal(ph)
}

func convertUserToModelUser(user *user) model.User {
	var gr model.Gender
	switch user.Gender {
	case genderMale:
		gr = model.GenderMale
	case genderFemale:
		gr = model.GenderFemale
	}

	return model.User{
		ID:                  user.ID,
		LastName:            user.LastName,
		FirstName:           user.FirstName,
		MiddleName:          user.MiddleName,
		Gender:              gr,
		DateOfBirth:         user.DateOfBirth,
		PlaceOfBirth:        user.PlaceOfBirth,
		Grade:               user.Grade,
		PhoneNumbers:        user.PhoneNumbers,
		Email:               user.Email,
		RegistrationAddress: user.RegistrationAddress,
		ResidentialAddress:  user.ResidentialAddress,
		Nationality:         user.Nationality,
		Insurance: model.Insurance{
			Number:  user.InsuranceNumber,
			HasScan: user.InsuranceHasScan,
		},
		TaxpayerNumber: user.TaxpayerNumber,
		Position:       user.Position,
		Department:     user.Department,
	}
}

type listUser struct {
	user
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
	VisasCount uint         `db:"visas_count"`
	HasScan    bool         `db:"has_scan"`
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
		VisasCount: p.VisasCount,
		HasScan:    p.HasScan,
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
	PassportID    uint64                  `db:"passport_id"`
	Number        string                  `db:"number"`
	IssuedState   string                  `db:"issued_state"`
	ValidTo       time.Time               `db:"valid_to"`
	ValidFrom     time.Time               `db:"valid_from"`
	NumberEntries model.VisaNumberEntries `db:"number_entries"`
}

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

type scan struct {
	ID          uint64    `db:"id"`
	UserID      uint64    `db:"user_id"`
	Type        scanType  `db:"type"`
	DocumentID  uint64    `db:"document_id"`
	Description string    `db:"description"`
	UploadedAt  time.Time `db:"created_at"`
}

type scanType string

const (
	scanTypePassport   scanType = "Паспорт"
	scanTypeTaxpayer   scanType = "ИНН"
	scanTypeInsurance  scanType = "СНИЛС"
	scanTypeContract   scanType = "Трудовой договор"
	scanTypePDP        scanType = "Согласие на обработку данных"
	scanTypeMilitary   scanType = "Военный билет"
	scanTypeEducation  scanType = "Документ об образовании"
	scanTypeTraining   scanType = "Сертификат"
	scanTypeBriefing   scanType = "Инструктаж"
	scanTypeWorkPermit scanType = "Разрешение на работу"
	scanTypeMarriage   scanType = "Свидетельство о браке"
	scanTypeBabyBirth  scanType = "Свидетельство о рождении"
	scanTypeOther      scanType = "Другое"
)

func convertScanToModelScan(s scan) model.Scan {
	var st model.ScanType
	switch s.Type {
	case scanTypePassport:
		st = model.ScanTypePassport
	case scanTypeTaxpayer:
		st = model.ScanTypeTaxpayer
	case scanTypeInsurance:
		st = model.ScanTypeInsurance
	case scanTypeContract:
		st = model.ScanTypeContract
	case scanTypePDP:
		st = model.ScanTypePDP
	case scanTypeMilitary:
		st = model.ScanTypeMilitary
	case scanTypeEducation:
		st = model.ScanTypeEducation
	case scanTypeTraining:
		st = model.ScanTypeTraining
	case scanTypeBriefing:
		st = model.ScanTypeBriefing
	case scanTypeWorkPermit:
		st = model.ScanTypeWorkPermit
	case scanTypeMarriage:
		st = model.ScanTypeMarriage
	case scanTypeBabyBirth:
		st = model.ScanTypeBabyBirth
	case scanTypeOther:
		st = model.ScanTypeOther
	}

	return model.Scan{
		ID:          s.ID,
		Type:        st,
		DocumentID:  s.DocumentID,
		Description: s.Description,
		UploadedAt:  s.UploadedAt,
	}
}

func convertModelScanToScan(ms model.Scan) scan {
	var t scanType
	switch ms.Type {
	case model.ScanTypePassport:
		t = scanTypePassport
	case model.ScanTypeTaxpayer:
		t = scanTypeTaxpayer
	case model.ScanTypeInsurance:
		t = scanTypeInsurance
	case model.ScanTypeContract:
		t = scanTypeContract
	case model.ScanTypePDP:
		t = scanTypePDP
	case model.ScanTypeMilitary:
		t = scanTypeMilitary
	case model.ScanTypeEducation:
		t = scanTypeEducation
	case model.ScanTypeTraining:
		t = scanTypeTraining
	case model.ScanTypeBriefing:
		t = scanTypeBriefing
	case model.ScanTypeWorkPermit:
		t = scanTypeWorkPermit
	case model.ScanTypeMarriage:
		t = scanTypeMarriage
	case model.ScanTypeBabyBirth:
		t = scanTypeBabyBirth
	case model.ScanTypeOther:
		t = scanTypeOther
	}

	return scan{
		ID:          ms.ID,
		DocumentID:  ms.DocumentID,
		Type:        t,
		Description: ms.Description,
	}
}
