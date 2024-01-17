package convert

import (
	"github.com/oapi-codegen/runtime/types"

	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
)

func FromAPIAddPassportRequest(req api.AddPassportJSONRequestBody) model.Passport {
	mp := model.Passport{
		IssuedBy:   req.IssuedBy,
		IssuedDate: req.IssuedDate.Time,
		Number:     req.Number,
	}
	switch req.Type {
	case api.Internal:
		mp.Type = model.PassportTypeInternal
	case api.External:
		mp.Type = model.PassportTypeExternal
	case api.Foreigners:
		mp.Type = model.PassportTypeForeigners
	}
	return mp
}

func FromAPIPutPassportRequest(passportID uint64, req api.PutPassportJSONRequestBody) model.Passport {
	mp := model.Passport{
		IssuedBy:   req.IssuedBy,
		IssuedDate: req.IssuedDate.Time,
		Number:     req.Number,
		ID:         passportID,
	}
	switch req.Type {
	case api.Internal:
		mp.Type = model.PassportTypeInternal
	case api.External:
		mp.Type = model.PassportTypeExternal
	case api.Foreigners:
		mp.Type = model.PassportTypeForeigners
	}
	return mp
}

func ToAPIGetPassportResponse(mp *model.Passport) api.GetPassportResponse {
	resp := api.GetPassportResponse{
		ID:         mp.ID,
		IssuedBy:   mp.IssuedBy,
		IssuedDate: types.Date{Time: mp.IssuedDate},
		Number:     mp.Number,
	}
	switch mp.Type {
	case model.PassportTypeInternal:
		resp.Type = api.Internal
	case model.PassportTypeExternal:
		resp.Type = api.External
	case model.PassportTypeForeigners:
		resp.Type = api.Foreigners
	}
	return resp
}

func ToAPIListPassports(eds []model.Passport) api.ListPassportsResponse {
	res := make([]api.Passport, len(eds))
	for i := 0; i < len(eds); i++ {
		res[i] = toAPIPassport(eds[i])
	}
	return res
}

func toAPIPassport(mp model.Passport) api.Passport {
	resp := api.GetPassportResponse{
		ID:         mp.ID,
		IssuedBy:   mp.IssuedBy,
		IssuedDate: types.Date{Time: mp.IssuedDate},
		Number:     mp.Number,
	}
	switch mp.Type {
	case model.PassportTypeInternal:
		resp.Type = api.Internal
	case model.PassportTypeExternal:
		resp.Type = api.External
	case model.PassportTypeForeigners:
		resp.Type = api.Foreigners
	}
	return resp
}

func ToAPIPassports(psps []model.Passport) []api.Passport {
	res := make([]api.Passport, len(psps))
	for i := 0; i < len(psps); i++ {
		res[i] = toAPIPassport(psps[i])
	}
	return res
}
