package errors

import (
	"errors"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/api"
	"github.com/Employee-s-file-cabinet/backend/internal/delivery/http/internal/response"
)

var (
	ErrInvalidContentType     = errors.New("invalid content type")
	ErrLimitRequestBodySize   = errors.New("request body too large")
	ErrBadContentLengthHeader = errors.New("bad content length header: missing or not a number")
	ErrInternalServerError    = errors.New("the server encountered a problem and could not process your request")
	ErrNotFoundRoute          = errors.New("the requested resource could not be found")
	ErrMethodNotAllowed       = errors.New("the method is not supported for this resource")
	ErrLoginFailure           = errors.New("login failed: invalid user ID or password")
)

// ReportError logs the server error, with or without stack trace.
func ReportError(r *http.Request, err error, withStack bool) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	if withStack {
		slog.Error(message, requestAttrs, "trace", trace)
	} else {
		slog.Error(message, requestAttrs)
	}
}

// ErrorMessage converts an error to api.Error and writes this one in JSON format to response writer
func ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, api.Error{Message: message}, headers)
	if err != nil {
		ReportError(r, err, false)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	ErrorMessage(w, r, http.StatusNotFound, ErrNotFoundRoute.Error(), nil)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	ErrorMessage(w, r, http.StatusMethodNotAllowed, ErrMethodNotAllowed.Error(), nil)
}
