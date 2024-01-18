package handlers

import (
	"net/http"

	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	srverr "github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/errors"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/response"
)

// @Success 200
// @Router  /health [get]
func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	status := h.statusService.HealthCheck(ctx)
	if err := response.JSON(w, http.StatusOK, api.Status(status)); err != nil {
		srverr.LogError(r, err, false)
		srverr.ResponseError(w, r,
			http.StatusInternalServerError,
			srverr.ErrInternalServerErrorMsg)
	}
}
