package convert

import (
	"github.com/oapi-codegen/runtime/types"

	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
)

func FromAPIAddContractRequest(req api.AddContractJSONRequestBody) model.Contract {
	mc := model.Contract{
		Number:          req.Number,
		WorkTypeID:      req.WorkTypeID,
		ProbationPeriod: req.ProbationPeriod,
		DateBegin:       req.DateFrom.Time,
	}

	if req.DateTo != nil {
		mc.DateEnd = &req.DateTo.Time
	}
	switch req.Type {
	case api.Permanent:
		mc.ContractType = model.ContractTypePermanent
	case api.Temporary:
		mc.ContractType = model.ContractTypeTemporary
	}
	return mc
}

func FromAPIPutContractRequest(contractID uint64, req api.PutContractJSONRequestBody) model.Contract {
	mc := model.Contract{
		ID:              contractID,
		Number:          req.Number,
		WorkTypeID:      req.WorkTypeID,
		ProbationPeriod: req.ProbationPeriod,
		DateBegin:       req.DateFrom.Time,
	}

	if req.DateTo != nil {
		mc.DateEnd = &req.DateTo.Time
	}
	switch req.Type {
	case api.Permanent:
		mc.ContractType = model.ContractTypePermanent
	case api.Temporary:
		mc.ContractType = model.ContractTypeTemporary
	}
	return mc
}

func ToAPIGetContractResponse(med *model.Contract) api.GetContractResponse {
	return toAPIContract(*med)
}

func ToAPIListContracts(eds []model.Contract) api.ListContractsResponse {
	res := make([]api.Contract, len(eds))
	for i := 0; i < len(eds); i++ {
		res[i] = toAPIContract(eds[i])
	}
	return res
}

func toAPIContract(med model.Contract) api.Contract {
	resp := api.GetContractResponse{
		ID:              med.ID,
		Number:          med.Number,
		WorkTypeID:      med.WorkTypeID,
		ProbationPeriod: med.ProbationPeriod,
		DateFrom:        types.Date{Time: med.DateBegin},
	}
	if med.DateEnd != nil {
		resp.DateTo = &types.Date{Time: *med.DateEnd}
	}
	switch med.ContractType {
	case model.ContractTypePermanent:
		resp.Type = api.Permanent
	case model.ContractTypeTemporary:
		resp.Type = api.Temporary
	}
	return resp
}
