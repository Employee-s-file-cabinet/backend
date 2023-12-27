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
		json.Unmarshal(v, &ph)
		return nil
	case string:
		json.Unmarshal([]byte(v), &ph)
		return nil
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
		InsuranceNumber:     user.InsuranceNumber,
		TaxpayerNumber:      user.TaxpayerNumber,
		Position:            user.Position,
		Department:          user.Department,
	}
}

type listUser struct {
	user
	TotalCount int `db:"total_count"`
}

func converListUserToModelUser(lu *listUser) model.User {
	return convertUserToModelUser(&lu.user)
}
