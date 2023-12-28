//nolint:all // code generated by oapi-codegen mostly
package api

import (
	"context"
	"time"

	vld "github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	otypes "github.com/oapi-codegen/runtime/types"
)

type contextKeyType string

const (
	BearerAuthScopes contextKeyType = "bearerAuth.Scopes"
)

// ContractType represents employee contract type.
type ContractType string

const (
	ContractTypePermanent    ContractType = "permanent"
	ContractTypeTemporary    ContractType = "temporary"
	ContractTypeSelfEmployed ContractType = "self_employed"
)

// Gender represents user gender.
type Gender string

const (
	GenderFemale Gender = "female"
	GenderMale   Gender = "male"
)

// ScanType represents document scan type.
type ScanType string

const (
	ScanTypeBabyBirth              ScanType = "baby_birth"
	ScanTypeBriefing               ScanType = "briefing"
	ScanTypeContract               ScanType = "contract"
	ScanTypeEducation              ScanType = "education"
	ScanTypeInsurance              ScanType = "insurance"
	ScanTypeMarriage               ScanType = "marriage"
	ScanTypeMilitary               ScanType = "military"
	ScanTypeOther                  ScanType = "other"
	ScanTypePassport               ScanType = "passport"
	ScanTypePersonalDataProcessing ScanType = "personal_data_processing"
	ScanTypeTaxpayer               ScanType = "taxpayer"
	ScanTypeTraining               ScanType = "training"
	ScanTypeWorkPermit             ScanType = "work_permit"
)

// VisaNumberEntries represents visa number entries.
type VisaNumberEntries string

const (
	VisaNumberEntriesMult VisaNumberEntries = "mult"
	VisaNumberEntriesN1   VisaNumberEntries = "1"
	VisaNumberEntriesN2   VisaNumberEntries = "2"
)

// WorkingModel represents user working model.
type WorkingModel string

const (
	WorkingModelHybrid   WorkingModel = "hybrid"
	WorkingModelInOffice WorkingModel = "in-office"
	WorkingModelRemote   WorkingModel = "remote"
)

const (
	ListUsersParamsSortByAlphabet   ListUsersParamsSortBy = "alphabet"
	ListUsersParamsSortByDepartment ListUsersParamsSortBy = "department"
)

// PassportType represents user passport type.
type PassportType string

const (
	PassportTypeExternal   PassportType = "external"
	PassportTypeForeigners PassportType = "foreigners"
	PassportTypeInternal   PassportType = "internal"
)

// Auth represents employee auth data: login and password.
type Auth struct {
	// Login employee login (email)
	Login string `json:"login"`

	// Password employee password
	Password string `json:"password"`
}

func (a Auth) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("login", string(a.Login), it.HasLengthBetween(5, 50), it.IsEmail()),
		vld.StringProperty("password", a.Password, it.HasLengthBetween(8, 15)),
	)
}

// Contract represents employee contract.
type Contract struct {
	DateFrom otypes.Date  `json:"date_from"`
	DateTo   *otypes.Date `json:"date_to,omitempty"`
	HasScan  bool         `json:"has_scan,omitempty"`
	ID       *uint64      `json:"id,omitempty"`
	Number   string       `json:"number"`
	Type     ContractType `json:"type"`
}

func (c Contract) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("number", c.Number, it.HasLengthBetween(2, 50)),
		vld.ComparableProperty[ContractType]("type",
			c.Type,
			it.IsOneOf[ContractType](ContractTypePermanent, ContractTypeTemporary, ContractTypeSelfEmployed)),
	)
}

// Department represents company department.
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	// RecruitedUsersNumber number of recruited employees
	RecruitedUsersNumber *int `json:"recruited_users_number,omitempty"`

	// UsersNumber current number of working employees
	UsersNumber int `json:"users_number"`
}

// Education represents employee education.
type Education struct {
	// DateFrom date of commencement of studies
	DateFrom otypes.Date `json:"date_from"`

	// DateTo date of graduation
	DateTo            *otypes.Date `json:"date_to,omitempty"`
	HasScan           bool         `json:"has_scan,omitempty"`
	ID                *uint64      `json:"id,omitempty"`
	IssuedInstitution string       `json:"issued_institution"`
	Number            string       `json:"number"`
	Program           string       `json:"program"`
}

func (e Education) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("number", e.Number, it.HasLengthBetween(2, 50)),
		vld.StringProperty("issued_institution", e.IssuedInstitution, it.HasLengthBetween(2, 150)),
		vld.StringProperty("program", e.Program, it.HasLengthBetween(2, 150)),
	)
}

// Error represents api error returned by server.
type Error struct {
	Code    *int   `json:"code,omitempty"`
	Message string `json:"message"`
}

type (
	BadRequestError   = Error
	UnauthorizedError = Error
)

// ShortUser represents short version of user data.
type ShortUser struct {
	Department   string                 `json:"department"`
	Email        string                 `json:"email"`
	FirstName    string                 `json:"first_name"`
	ID           *uint64                `json:"id,omitempty"`
	LastName     string                 `json:"last_name"`
	MiddleName   string                 `json:"middle_name,omitempty"`
	PhoneNumbers map[string]PhoneNumber `json:"phone_numbers,omitempty"`
	Position     string                 `json:"position"`
}

// FullUser represents full version of user data.
type FullUser struct {
	ShortUser
	DateOfBirth            otypes.Date             `json:"date_of_birth"`
	Finance                *UserFinance            `json:"finance,omitempty"`
	ForeignLanguages       []string                `json:"foreign_languages,omitempty"`
	Gender                 Gender                  `json:"gender"`
	Grade                  string                  `json:"grade"`
	Insurance              Insurance               `json:"insurance"`
	Military               *Military               `json:"military,omitempty"`
	Nationality            string                  `json:"nationality"`
	PersonalDataProcessing *PersonalDataProcessing `json:"personal_data_processing,omitempty"`
	PlaceOfBirth           string                  `json:"place_of_birth"`
	PositionTrack          []PositionTrackItem     `json:"position_track"`
	RegistrationAddress    string                  `json:"registration_address"`
	ResidentialAddress     string                  `json:"residential_address"`
	Taxpayer               Taxpayer                `json:"taxpayer"`
	WorkPermit             *WorkPermit             `json:"work_permit,omitempty"`
	WorkingModel           *WorkingModel           `json:"working_model,omitempty"`
}

func (u FullUser) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("first_name", u.FirstName, it.HasLengthBetween(2, 150)),
		vld.StringProperty("last_name", u.LastName, it.HasLengthBetween(2, 150)),
		vld.StringProperty("middle_name", u.MiddleName, it.HasLengthBetween(2, 150)).
			When(u.MiddleName != ""),
		vld.StringProperty("position", u.Position, it.HasLengthBetween(2, 150)),
		vld.StringProperty("department", u.Department, it.HasLengthBetween(2, 150)),
		vld.StringProperty("email", u.Email, it.HasLengthBetween(5, 50), it.IsEmail()),
		vld.ValidMapProperty[PhoneNumber]("phone_numbers", u.PhoneNumbers),
		vld.StringProperty("grade", u.Grade, it.HasExactLength(1)),
		vld.When(u.WorkingModel != nil).
			At(vld.PropertyName("working_model")).
			Then(vld.NilComparable(u.WorkingModel, it.IsOneOf[WorkingModel](
				WorkingModelHybrid,
				WorkingModelInOffice,
				WorkingModelRemote))),
		vld.ComparableProperty[Gender]("gender",
			u.Gender,
			it.IsOneOf[Gender](GenderMale, GenderFemale)),
		vld.StringProperty("place_of_birth", u.PlaceOfBirth, it.HasLengthBetween(2, 150)),
		vld.StringProperty("registration_address", u.RegistrationAddress, it.HasLengthBetween(2, 150)),
		vld.StringProperty("residential_address", u.ResidentialAddress, it.HasLengthBetween(2, 150)),
		vld.StringProperty("nationality", u.Nationality, it.HasLengthBetween(2, 150)),
		vld.EachStringProperty("foreign_languages", u.ForeignLanguages, it.HasLengthBetween(2, 50)),
		vld.ValidSliceProperty[PositionTrackItem]("position_track", u.PositionTrack),
		vld.When(u.Military != nil).
			At(vld.PropertyName("military")).
			Then(vld.ValidProperty("military", u.Military)),
		vld.ValidProperty("insurance", u.Insurance),
		vld.ValidProperty("taxpayer", u.Taxpayer),
		vld.When(u.WorkPermit != nil).
			At(vld.PropertyName("work_permit")).
			Then(vld.ValidProperty("work_permit", u.WorkPermit)),
	)
}

// Insurance represents employee insurance document data.
type Insurance struct {
	HasScan bool   `json:"has_scan,omitempty"`
	Number  string `json:"number"`
}

func (i Insurance) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("number",
			i.Number,
			it.HasExactLength(11),
			consistOnlyNumbersFormat(),
			hasCorrectInsuranceChecksum()),
	)
}

// Military represents employee military document data.
type Military struct {
	Category    string `json:"category"`
	Comissariat string `json:"comissariat"`
	HasScan     bool   `json:"has_scan,omitempty"`
	Rank        string `json:"rank"`
	Speciality  string `json:"speciality"`
}

func (m Military) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("category", m.Category, it.HasLengthBetween(1, 2)),
		vld.StringProperty("comissariat", m.Comissariat, it.HasLengthBetween(2, 150)),
		vld.StringProperty("rank", m.Rank, it.HasLengthBetween(2, 150)),
		vld.StringProperty("speciality", m.Speciality, it.HasLengthBetween(6, 7)),
	)
}

// InitChangePasswordRequest represents a first user request
// to change the password.
type InitChangePasswordRequest struct {
	// Login employee login (email)
	Login string `json:"login"`
}

func (ir InitChangePasswordRequest) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("login", ir.Login, it.HasLengthBetween(5, 50), it.IsEmail()),
	)
}

// ChangePasswordRequest represents a final user request
// to change the password.
type ChangePasswordRequest struct {
	// Key a special key sent to the employee’s email
	Key string `json:"key"`

	// Password a new employee password
	Password string `json:"password"`
}

func (cr ChangePasswordRequest) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("key", cr.Key, it.HasExactLength(36)),
		vld.StringProperty("key", cr.Key, it.Matches(rxBase64)),
		vld.StringProperty("password", cr.Password, it.HasLengthBetween(8, 15)),
	)
}

// Passport represents employee passport data.
type Passport struct {
	HasScan    bool         `json:"has_scan,omitempty"`
	ID         *uint64      `json:"id,omitempty"`
	IssuedBy   string       `json:"issued_by"`
	IssuedDate otypes.Date  `json:"issued_date"`
	Number     string       `json:"number"`
	Type       PassportType `json:"type"`
}

func (p Passport) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("issued_by", p.IssuedBy, it.HasLengthBetween(2, 150)),
		vld.StringProperty("number",
			p.Number,
			it.HasLengthBetween(2, 50),
			consistOnlyNumbersFormat()),
		vld.ComparableProperty[PassportType]("type",
			p.Type,
			it.IsOneOf[PassportType](
				PassportTypeExternal,
				PassportTypeInternal,
				PassportTypeForeigners)),
	)
}

// PersonalDataProcessing represents employee personal
// data processing document data.
type PersonalDataProcessing struct {
	HasScan bool `json:"has_scan,omitempty"`
}

// PhoneNumber represents employee phone number.
type PhoneNumber string

func (pn PhoneNumber) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("phone_number", string(pn), it.HasLengthBetween(2, 15)),
	)
}

// PositionTrackItem represents one item in employee position track in the company.
type PositionTrackItem struct {
	DateFrom otypes.Date  `json:"date_from"`
	DateTo   *otypes.Date `json:"date_to,omitempty"`
	Position string       `json:"position"`
}

func (pt PositionTrackItem) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("position", pt.Position, it.HasLengthBetween(2, 150)),
	)
}

// Scan represents employee document scan.
type Scan struct {
	Description string    `json:"description,omitempty"`
	DocumentID  *int      `json:"document_id,omitempty"`
	ID          *uint64   `json:"id,omitempty"`
	Type        ScanType  `json:"type"`
	UploadAt    time.Time `json:"upload_at,omitempty"`
}

// Taxpayer represents employee tax document data.
type Taxpayer struct {
	HasScan bool   `json:"has_scan,omitempty"`
	Number  string `json:"number"`
}

func (tp Taxpayer) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("number",
			tp.Number,
			it.HasLengthBetween(10, 12),
			consistOnlyNumbersFormat(),
			hasCorrectTaxpayerChecksum()),
	)
}

type Money uint64

// Training represents employee training (course, seminar, etc).
type Training struct {
	// Cost training, in their minor unit form
	Cost *Money `json:"cost,omitempty"`

	DateFrom          otypes.Date  `json:"date_from"`
	DateTo            *otypes.Date `json:"date_to,omitempty"`
	HasScan           bool         `json:"has_scan,omitempty"`
	ID                *uint64      `json:"id,omitempty"`
	IssuedInstitution string       `json:"issued_institution"`
	Number            string       `json:"number,omitempty"`
	Program           string       `json:"program"`
}

func (t Training) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("issued_institution", t.IssuedInstitution, it.HasLengthBetween(2, 150)),
		vld.StringProperty("program", t.Program, it.HasLengthBetween(2, 150)),
		vld.StringProperty("number", t.Number, it.HasLengthBetween(2, 50)),
	)
}

// UserFinance represents employee finance data.
// The data imported from the external finance API,
// cannot be changed by the current API user.
type UserFinance struct {
	// IncomeTax tax paid to revenue service per month, in their minor unit form
	IncomeTax *Money `json:"income_tax"`

	// Salary gross salary per month, in their minor unit form
	Salary *Money `json:"salary"`

	// SocialSecurityTax tax paid to social services per month, in their minor unit form
	SocialSecurityTax *Money `json:"social_security_tax"`
}

// Vacation represents employee vacation.
type Vacation struct {
	DateFrom otypes.Date `json:"date_from"`
	DateTo   otypes.Date `json:"date_to"`
	ID       *uint64     `json:"id,omitempty"`
}

func (v Vacation) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(ctx)
}

// Visa represents employee visa.
type Visa struct {
	HasScan       bool              `json:"has_scan,omitempty"`
	ID            *uint64           `json:"id,omitempty"`
	IssuedState   string            `json:"issued_state"`
	Number        string            `json:"number"` // can have letter characters (ex: the USA visa)
	NumberEntries VisaNumberEntries `json:"number_entries"`
	ValidFrom     otypes.Date       `json:"valid_from"`
	ValidTo       otypes.Date       `json:"valid_to"`
}

func (v Visa) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("issued_state", v.IssuedState, it.HasLengthBetween(2, 50)),
		vld.StringProperty("number", v.Number, it.HasLengthBetween(2, 50)),
		vld.ComparableProperty[VisaNumberEntries]("number_entries",
			v.NumberEntries,
			it.IsOneOf[VisaNumberEntries](
				VisaNumberEntriesN1,
				VisaNumberEntriesN2,
				VisaNumberEntriesMult)),
	)
}

// WorkPermit represents employee work permit.
type WorkPermit struct {
	HasScan   bool         `json:"has_scan,omitempty"`
	Number    string       `json:"number"`
	ValidFrom *otypes.Date `json:"valid_from,omitempty"`
	ValidTo   otypes.Date  `json:"valid_to"`
}

func (wp WorkPermit) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("number",
			wp.Number,
			it.HasLengthBetween(2, 50),
			consistOnlyNumbersFormat()), // TODO: unknown format, length
	)
}

// Token represents a token to access the API.
type Token struct {
	// The access token issued by the server
	AccessToken string `json:"access_token"`

	// The type of the token issued
	TokenType string `json:"token_type"`

	// The lifetime in seconds of the access token
	ExpiresIn int `json:"expires_in"`
}

// ------------------------------------------------------------------
// Parameters
// ------------------------------------------------------------------

type CheckKeyParams struct {
	// Key a special key sent to the employee’s email
	Key string `form:"key" json:"key"`
}

func (cp CheckKeyParams) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.StringProperty("key", cp.Key, it.HasExactLength(36)),
		vld.StringProperty("key", string(cp.Key), it.Matches(rxBase64)),
	)
}

type ListUsersParams struct {
	// Limit maximum number of results to return
	Limit *uint `form:"limit,omitempty" json:"limit,omitempty"`

	// Query query to find by
	Query *string `form:"query,omitempty" json:"query,omitempty"`

	// Page page number with results to return
	Page *uint `form:"page,omitempty" json:"page,omitempty"`

	// SortBy type of sort result by
	SortBy *ListUsersParamsSortBy `form:"sort_by,omitempty" json:"sort_by,omitempty"`
}

func (lp ListUsersParams) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(lp.SortBy != nil).
			At(vld.PropertyName("sort_by")).
			Then(vld.NilComparable(lp.SortBy,
				it.IsOneOf[ListUsersParamsSortBy](
					ListUsersParamsSortByAlphabet,
					ListUsersParamsSortByDepartment))),
		vld.When(lp.Query != nil).
			At(vld.PropertyName("query")).
			Then(vld.NilString(lp.Query, it.HasMaxLength(150))),
	)
}

type ListUsersParamsSortBy string

// GetUserParams defines parameters for GetUser.
type GetUserParams struct {
	// Expanded whether to return detailed data on passports, contracts, vacations, etc. along with user data (default - no)
	Expanded *bool `form:"expanded,omitempty" json:"expanded,omitempty"`
}

// ------------------------------------------------------------------
// Custom requests
// ------------------------------------------------------------------

type PatchFullUserJSONRequestBody struct {
	DateOfBirth         *otypes.Date           `json:"date_of_birth,omitempty"`
	Department          *string                `json:"department,omitempty"`
	Email               *string                `json:"email,omitempty"`
	FirstName           *string                `json:"first_name,omitempty"`
	ForeignLanguages    []string               `json:"foreign_languages,omitempty"`
	Gender              *Gender                `json:"gender,omitempty"`
	Grade               *string                `json:"grade,omitempty"`
	ID                  *uint64                `json:"id,omitempty"`
	Insurance           *Insurance             `json:"insurance,omitempty"`
	LastName            *string                `json:"last_name,omitempty"`
	MiddleName          *string                `json:"middle_name,omitempty"`
	Military            *Military              `json:"military,omitempty"`
	Nationality         *string                `json:"nationality,omitempty"`
	PhoneNumbers        map[string]PhoneNumber `json:"phone_numbers,omitempty"`
	PlaceOfBirth        *string                `json:"place_of_birth,omitempty"`
	Position            *string                `json:"position,omitempty"`
	PositionTrack       []PositionTrackItem    `json:"position_track,omitempty"`
	RegistrationAddress *string                `json:"registration_address,omitempty"`
	ResidentialAddress  *string                `json:"residential_address,omitempty"`
	Taxpayer            *Taxpayer              `json:"taxpayer,omitempty"`
	WorkPermit          *WorkPermit            `json:"work_permit,omitempty"`
	WorkingModel        *WorkingModel          `json:"working_model,omitempty"`
}

func (pu PatchFullUserJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pu.FirstName != nil).
			At(vld.PropertyName("first_name")).
			Then(vld.NilString(pu.FirstName, it.HasLengthBetween(2, 150))),
		vld.When(pu.LastName != nil).
			At(vld.PropertyName("last_name")).
			Then(vld.NilString(pu.LastName, it.HasLengthBetween(2, 150))),
		vld.When(pu.MiddleName != nil && *pu.MiddleName != "").
			At(vld.PropertyName("middle_name")).
			Then(vld.NilString(pu.MiddleName, it.HasLengthBetween(2, 150))),
		vld.When(pu.Position != nil && *pu.Position != "").
			At(vld.PropertyName("position")).
			Then(vld.NilString(pu.Position, it.HasLengthBetween(2, 150))),
		vld.When(pu.Department != nil && *pu.Department != "").
			At(vld.PropertyName("department")).
			Then(vld.NilString(pu.Department, it.HasLengthBetween(2, 150))),
		vld.When(pu.Email != nil && *pu.Email != "").
			At(vld.PropertyName("email")).
			Then(vld.NilString(pu.Email, it.HasLengthBetween(5, 50), it.IsEmail())),
		vld.ValidMapProperty[PhoneNumber]("phone_numbers", pu.PhoneNumbers),
		vld.When(pu.Grade != nil && *pu.Grade != "").
			At(vld.PropertyName("grade")).
			Then(vld.NilString(pu.Grade, it.HasExactLength(1))),
		vld.When(pu.WorkingModel != nil).
			At(vld.PropertyName("working_model")).
			Then(vld.NilComparable(pu.WorkingModel, it.IsOneOf[WorkingModel](
				WorkingModelHybrid,
				WorkingModelInOffice,
				WorkingModelRemote))),
		vld.When(pu.Gender != nil).
			At(vld.PropertyName("gender")).
			Then(vld.NilComparable(pu.Gender,
				it.IsOneOf[Gender](
					GenderMale,
					GenderFemale))),
		vld.When(pu.PlaceOfBirth != nil && *pu.PlaceOfBirth != "").
			At(vld.PropertyName("place_of_birth")).
			Then(vld.NilString(pu.PlaceOfBirth, it.HasLengthBetween(2, 150))),
		vld.When(pu.RegistrationAddress != nil && *pu.RegistrationAddress != "").
			At(vld.PropertyName("registration_address")).
			Then(vld.NilString(pu.RegistrationAddress, it.HasLengthBetween(2, 150))),
		vld.When(pu.ResidentialAddress != nil && *pu.ResidentialAddress != "").
			At(vld.PropertyName("residential_address")).
			Then(vld.NilString(pu.ResidentialAddress, it.HasLengthBetween(2, 150))),
		vld.When(pu.Nationality != nil && *pu.Nationality != "").
			At(vld.PropertyName("nationality")).
			Then(vld.NilString(pu.Nationality, it.HasLengthBetween(2, 150))),
		vld.EachStringProperty("foreign_languages", pu.ForeignLanguages, it.HasLengthBetween(2, 50)),
		vld.ValidSliceProperty[PositionTrackItem]("position_track", pu.PositionTrack),
		vld.When(pu.Military != nil).
			At(vld.PropertyName("military")).
			Then(vld.ValidProperty("military", pu.Military)),
		vld.When(pu.Insurance != nil).
			At(vld.PropertyName("insurance")).
			Then(vld.ValidProperty("insurance", pu.Insurance)),
		vld.When(pu.Taxpayer != nil).
			At(vld.PropertyName("taxpayer")).
			Then(vld.ValidProperty("taxpayer", pu.Taxpayer)),
		vld.When(pu.WorkPermit != nil).
			At(vld.PropertyName("work_permit")).
			Then(vld.ValidProperty("work_permit", pu.WorkPermit)),
	)
}

type PatchContractJSONRequestBody struct {
	DateFrom *otypes.Date  `json:"date_from,omitempty"`
	DateTo   *otypes.Date  `json:"date_to,omitempty"`
	HasScan  *bool         `json:"has_scan,omitempty"`
	ID       *uint64       `json:"id,omitempty"`
	Number   *string       `json:"number,omitempty"`
	Type     *ContractType `json:"type,omitempty"`
}

func (pc PatchContractJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pc.Number != nil).
			At(vld.PropertyName("number")).
			Then(vld.NilString(pc.Number, it.HasLengthBetween(2, 50))),
		vld.When(pc.Type != nil).
			At(vld.PropertyName("type")).
			Then(vld.NilComparable(pc.Type,
				it.IsOneOf[ContractType](
					ContractTypePermanent,
					ContractTypeTemporary,
					ContractTypeSelfEmployed))),
	)
}

type PatchEducationJSONRequestBody struct {
	// DateFrom date of commencement of studies
	DateFrom *otypes.Date `json:"date_from,omitempty"`

	// DateTo date of graduation
	DateTo            *otypes.Date `json:"date_to,omitempty"`
	HasScan           *bool        `json:"has_scan,omitempty"`
	ID                *uint64      `json:"id,omitempty"`
	IssuedInstitution *string      `json:"issued_institution,omitempty"`
	Number            *string      `json:"number,omitempty"`
	Program           *string      `json:"program,omitempty"`
}

func (pe PatchEducationJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pe.Number != nil).
			At(vld.PropertyName("number")).
			Then(vld.NilString(pe.Number, it.HasLengthBetween(2, 50))),
		vld.When(pe.IssuedInstitution != nil).
			At(vld.PropertyName("issued_institution")).
			Then(vld.NilString(pe.IssuedInstitution, it.HasLengthBetween(2, 150))),
		vld.When(pe.Program != nil).
			At(vld.PropertyName("program")).
			Then(vld.NilString(pe.Program, it.HasLengthBetween(2, 150))),
	)
}

type PatchPassportJSONRequestBody struct {
	HasScan    *bool         `json:"has_scan,omitempty"`
	ID         *uint64       `json:"id,omitempty"`
	IssuedBy   *string       `json:"issued_by,omitempty"`
	IssuedDate *string       `json:"issued_date,omitempty"`
	Number     *string       `json:"number,omitempty"`
	Type       *PassportType `json:"type,omitempty"`
}

func (pp PatchPassportJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pp.IssuedBy != nil).
			At(vld.PropertyName("issued_by")).
			Then(vld.NilString(pp.IssuedBy, it.HasLengthBetween(2, 150))),
		vld.When(pp.Number != nil).
			At(vld.PropertyName("number")).
			Then(vld.NilString(pp.Number,
				it.HasLengthBetween(2, 50),
				consistOnlyNumbersFormat())),
		vld.When(pp.Type != nil).
			At(vld.PropertyName("type")).
			Then(vld.NilComparable(pp.Type,
				it.IsOneOf[PassportType](
					PassportTypeExternal,
					PassportTypeInternal,
					PassportTypeForeigners))),
	)
}

type PatchVisaJSONRequestBody struct {
	HasScan       *bool              `json:"has_scan,omitempty"`
	ID            *uint64            `json:"id,omitempty"`
	IssuedState   *string            `json:"issued_state,omitempty"`
	Number        *string            `json:"number,omitempty"`
	NumberEntries *VisaNumberEntries `json:"number_entries,omitempty"`
	ValidFrom     *otypes.Date       `json:"valid_from,omitempty"`
	ValidTo       *otypes.Date       `json:"valid_to,omitempty"`
}

func (pv PatchVisaJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pv.IssuedState != nil).
			At(vld.PropertyName("issued_state")).
			Then(vld.NilString(pv.IssuedState, it.HasLengthBetween(2, 50))),
		vld.When(pv.Number != nil).
			At(vld.PropertyName("number")).
			Then(vld.NilString(pv.Number, it.HasLengthBetween(2, 50))),
		vld.When(pv.NumberEntries != nil).
			At(vld.PropertyName("number_entries")).
			Then(vld.NilComparable(pv.NumberEntries,
				it.IsOneOf[VisaNumberEntries](
					VisaNumberEntriesN1,
					VisaNumberEntriesN2,
					VisaNumberEntriesMult))),
	)
}

type UploadScanMultipartRequestBody struct {
	Description *string
	DocumentID  *int
	FileName    otypes.File
	Type        ScanType
}

func (us UploadScanMultipartRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(us.Description != nil && *us.Description != "").
			At(vld.PropertyName("description")).
			Then(vld.NilString(us.Description, it.HasLengthBetween(2, 150))),
		vld.ComparableProperty[ScanType]("type",
			us.Type,
			it.IsOneOf[ScanType](
				ScanTypeBabyBirth,
				ScanTypeBriefing,
				ScanTypeContract,
				ScanTypeEducation,
				ScanTypeInsurance,
				ScanTypeMarriage,
				ScanTypeMilitary,
				ScanTypeOther,
				ScanTypePassport,
				ScanTypePersonalDataProcessing,
				ScanTypeTaxpayer,
				ScanTypeTraining,
				ScanTypeWorkPermit)),
	)
}

type PatchTrainingJSONRequestBody struct {
	// Cost cost per person, in their minor unit form
	Cost *Money `json:"cost,omitempty"`

	// DateFrom start date of training
	DateFrom *otypes.Date `json:"date_from,omitempty"`

	// DateTo end date of training
	DateTo            *otypes.Date `json:"date_to,omitempty"`
	HasScan           *bool        `json:"has_scan,omitempty"`
	ID                *uint64      `json:"id,omitempty"`
	IssuedInstitution *string      `json:"issued_institution,omitempty"`
	Number            *string      `json:"number,omitempty"`
	Program           *string      `json:"program,omitempty"`
}

func (pt PatchTrainingJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(
		ctx,
		vld.When(pt.IssuedInstitution != nil).
			At(vld.PropertyName("issued_institution")).
			Then(vld.NilString(pt.IssuedInstitution, it.HasLengthBetween(2, 150))),
		vld.When(pt.Program != nil).
			At(vld.PropertyName("program")).
			Then(vld.NilString(pt.Program, it.HasLengthBetween(2, 150))),
		vld.When(pt.Number != nil).
			At(vld.PropertyName("number")).
			Then(vld.NilString(pt.Number, it.HasLengthBetween(2, 50))),
	)
}

type PatchVacationJSONRequestBody struct {
	DateFrom *otypes.Date `json:"date_from,omitempty"`
	DateTo   *otypes.Date `json:"date_to,omitempty"`
	ID       *uint64      `json:"id,omitempty"`
}

func (v PatchVacationJSONRequestBody) Validate(ctx context.Context, validator *vld.Validator) error {
	return validator.Validate(ctx)
}

// ------------------------------------------------------------------
// Custom response
// ------------------------------------------------------------------

type ListUsersJSONResponseBody struct {
	Users       []ShortUser `json:"users"`
	TotalUsers  int         `json:"total_users"`
	TotalPages  int         `json:"total_pages"`
	CurrentPage int         `json:"current_page"`
}

type GetUserJSONResponseBody struct {
	FullUser
	Contracts  []Contract  `json:"contracts"`
	Educations []Education `json:"educations"`
	Passports  []Passport  `json:"passports"`
	Trainings  []Training  `json:"trainings"`
	Vacations  []Vacation  `json:"vacations"`
}