package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/muonsoft/validation/validator"

	serr "github.com/Employee-s-file-cabinet/backend/internal/delivery/http/errors"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/convert"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/request"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/response"
	"github.com/Employee-s-file-cabinet/backend/internal/service/user"
)

// @Produce application/json
// @Success 200 {object} api.ListPassportsResponse
// @Router  /users/{user_id}/passports [get]
func (h *handler) ListPassports(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := r.Context()

	psps, err := h.userService.ListPassports(ctx, userID)
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

	if err := response.JSON(w, http.StatusOK, convert.ToAPIListPassports(psps)); err != nil {
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
	}
}

// @Accept  application/json
// @Param   body body api.AddPassportJSONRequestBody true ""
// @Failure 409  {object} api.Error "passport already exists"
// @Router  /users/{user_id}/passports [post]
func (h *handler) AddPassport(w http.ResponseWriter, r *http.Request, userID uint64) {
	ctx := r.Context()

	var p api.AddPassportJSONRequestBody
	if err := request.DecodeJSONStrict(w, r, &p); err != nil {
		serr.ErrorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := p.Validate(ctx, validator.Instance()); err != nil {
		msg := api.ValidationErrorMessage(err)
		serr.ErrorMessage(w, r, http.StatusBadRequest, msg, nil)
		return
	}

	id, err := h.userService.AddPassport(ctx, userID, convert.FromAPIAddPassportRequest(p))
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
			"/passports/"+strconv.FormatUint(id, 10))
}

// @Router /users/{user_id}/passports/{passport_id} [delete]
func (h *handler) DeletePassport(w http.ResponseWriter, r *http.Request, userID uint64, passportID uint64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// @Produce application/json
// @Success 200 {object} api.GetPassportResponse
// @Router  /users/{user_id}/passports/{passport_id} [get]
func (h *handler) GetPassport(w http.ResponseWriter, r *http.Request, userID, passportID uint64, params api.GetPassportParams) {
	ctx := r.Context()

	if params.Expanded != nil && *params.Expanded {
		// TODO: implement this
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	p, err := h.userService.GetPassport(ctx, userID, passportID)
	if err != nil {
		if errors.Is(err, user.ErrPassportNotFound) {
			serr.ErrorMessage(w, r, http.StatusNotFound, user.ErrPassportNotFound.Error(), nil)
			return
		}
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
		return
	}

	if err := response.JSON(w, http.StatusOK, convert.ToAPIGetPassportResponse(p)); err != nil {
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
	}
}

// @Accept application/json
// @Param   body body api.PatchPassportJSONRequestBody true ""
// @Router  /users/{user_id}/passports/{passport_id} [patch]
func (h *handler) PatchPassport(w http.ResponseWriter, r *http.Request, userID uint64, passportID uint64) {
	ctx := r.Context()

	var patch api.PatchPassportJSONRequestBody
	// TODO: decode patch from request body

	if err := patch.Validate(ctx, validator.Instance()); err != nil {
		msg := api.ValidationErrorMessage(err)
		serr.ErrorMessage(w, r, http.StatusBadRequest, msg, nil)
		return
	}

	w.WriteHeader(http.StatusNotImplemented)
}

// @Accept  application/json
// @Param   body body api.PutPassportJSONRequestBody true ""
// @Router  /users/{user_id}/passports/{passport_id} [put]
func (h *handler) PutPassport(w http.ResponseWriter, r *http.Request, userID, passportID uint64) {
	ctx := r.Context()

	var p api.PutPassportJSONRequestBody
	if err := request.DecodeJSONStrict(w, r, &p); err != nil {
		serr.ErrorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := p.Validate(ctx, validator.Instance()); err != nil {
		msg := api.ValidationErrorMessage(err)
		serr.ErrorMessage(w, r, http.StatusBadRequest, msg, nil)
		return
	}

	err := h.userService.UpdatePassport(ctx, userID, convert.FromAPIPutPassportRequest(passportID, p))
	if err != nil {
		if errors.Is(err, user.ErrPassportNotFound) {
			serr.ErrorMessage(w, r, http.StatusNotFound, user.ErrPassportNotFound.Error(), nil)
			return
		}
		serr.ReportError(r, err, false)
		serr.ErrorMessage(w, r,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
			nil)
		return
	}
}
