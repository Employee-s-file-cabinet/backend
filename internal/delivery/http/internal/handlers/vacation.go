package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/muonsoft/validation/validator"
	"github.com/oapi-codegen/runtime/types"

	serr "github.com/Employee-s-file-cabinet/backend/internal/delivery/http/errors"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/request"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/response"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user/model"
)

// @Produce application/json
// @Success 200 {array} api.Vacation
// @Router  /users/{user_id}/vacations [get]
func (h *handler) ListVacations(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := r.Context()

	vacations, err := h.userService.ListVacations(ctx, userID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			serr.ErrorMessage(w, r, http.StatusNotFound, user.ErrUserNotFound.Error(), nil)
			return
		}
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
		return
	}

	if err := response.JSON(w, http.StatusOK, convertVacationsToAPIVacations(vacations)); err != nil {
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
	}
}

// @Accept application/json
// @Param   body body api.Vacation true ""
// @Failure 409  {object} api.Error "vacation already exists"
// @Router  /users/{user_id}/vacations [post]
func (h *handler) AddVacation(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := r.Context()

	var v api.Vacation
	if err := request.DecodeJSONStrict(w, r, &v); err != nil {
		serr.ErrorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := v.Validate(ctx, validator.Instance()); err != nil {
		msg := api.ValidationErrorMessage(err)
		serr.ErrorMessage(w, r, http.StatusBadRequest, msg, nil)
		return
	}

	id, err := h.userService.AddVacation(ctx, userID, convertAPIVacationToVacation(v))
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			serr.ErrorMessage(w, r, http.StatusConflict, user.ErrUserNotFound.Error(), nil)
			return
		}
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
		return
	}

	w.Header().Set("Location",
		api.BaseURL+"/users/"+strconv.FormatUint(userID, 10)+
			"/vacations/"+strconv.FormatUint(id, 10))
}

// @Router /users/{user_id}/vacations/{vacation_id} [delete]
func (h *handler) DeleteVacation(w http.ResponseWriter, r *http.Request, userID uint64, vacationID uint64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// @Accept application/json
// @Success 200 {object} api.Vacation
// @Router  /users/{user_id}/vacations/{vacation_id} [get]
func (h *handler) GetVacation(w http.ResponseWriter, r *http.Request, userID uint64, vacationID uint64) {
	ctx := r.Context()

	v, err := h.userService.GetVacation(ctx, userID, vacationID)
	if err != nil {
		if errors.Is(err, user.ErrVacationNotFound) {
			serr.ErrorMessage(w, r, http.StatusNotFound, user.ErrVacationNotFound.Error(), nil)
			return
		}
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
		return
	}

	if err := response.JSON(w, http.StatusOK, convertVacationToAPIVacation(v)); err != nil {
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
	}
}

// @Accept application/json
// @Param   body body api.PatchVacationJSONRequestBody true ""
// @Router  /users/{user_id}/vacations/{vacation_id} [patch]
func (h *handler) PatchVacation(w http.ResponseWriter, r *http.Request, userID uint64, vacationID uint64) {
	ctx := r.Context()

	var patch api.PatchVacationJSONRequestBody
	// TODO: decode patch from request body

	if err := patch.Validate(ctx, validator.Instance()); err != nil {
		msg := api.ValidationErrorMessage(err)
		serr.ErrorMessage(w, r, http.StatusBadRequest, msg, nil)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

func convertVacationsToAPIVacations(vcs []model.Vacation) []api.Vacation {
	res := make([]api.Vacation, len(vcs))
	for i := 0; i < len(vcs); i++ {
		res[i] = convertVacationToAPIVacation(&vcs[i])
	}
	return res
}

func convertVacationToAPIVacation(mv *model.Vacation) api.Vacation {
	return api.Vacation{
		ID:       &mv.ID,
		DateFrom: types.Date{Time: mv.DateBegin},
		DateTo:   types.Date{Time: mv.DateEnd},
	}
}

func convertAPIVacationToVacation(tr api.Vacation) model.Vacation {
	return model.Vacation{
		DateBegin: tr.DateFrom.Time,
		DateEnd:   tr.DateTo.Time,
	}
}
