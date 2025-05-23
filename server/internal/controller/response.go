package controller

import (
	"encoding/json"
	"media/internal/errors"
	"media/internal/logger"
	"net/http"

	"github.com/rs/zerolog"
)

var log = zerolog.New(logger.Output()).With().
	Timestamp().
	Logger()

type HTTPError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func SendError(w http.ResponseWriter, httpCode int, err error, code string) {
	errorStruct := HTTPError{
		Message: err.Error(),
		Code:    code,
	}
	if e, ok := err.(*errors.PressError); ok {
		errorStruct.Message = e.Public()
	}
	log.Error().Stack().Err(err).Msg("request error")

	w.WriteHeader(httpCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(errorStruct)
}

func SendJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		InternalError(w, err)
	}

	w.WriteHeader(http.StatusOK)
}

func InternalError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusInternalServerError, errors.New("Something went wrong").WithInternal(err.Error()), "Internal Error")
}

func UnauthorizedError(w http.ResponseWriter) {
	SendError(w, http.StatusUnauthorized, errors.New("not authorized"), "Not authorized")
}

func ValidationError(w http.ResponseWriter, err error) {
	SendError(w, http.StatusUnprocessableEntity, err, "Validation Error")
}
