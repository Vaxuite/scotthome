package routing

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type HTTPError struct {
	Reason string `json:"reason"`
	Error  string `json:"error"`
}

func NewHTTPError(reason string, errorString string) *HTTPError {
	return &HTTPError{Reason: reason, Error: errorString}
}

func RespondWithHTTPError(errorDescription string, err error, w http.ResponseWriter) {
	zap.S().Errorw(errorDescription,
		"error", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(NewHTTPError(errorDescription, err.Error()).marshallError())
	if err != nil {
		zap.S().Fatalw("Error writing response",
			"err", err.Error())
	}
}

func RespondWithHTTPWarn(errorDescription string, err error, w http.ResponseWriter) {
	zap.S().Warnw(errorDescription,
		"error", err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	_, err = w.Write(NewHTTPError(errorDescription, err.Error()).marshallError())
	if err != nil {
		zap.S().Fatalw("Error writing response",
			"err", err.Error())
	}
}

func (httpError *HTTPError) marshallError() []byte {
	payload, err := json.Marshal(httpError)
	if err != nil {
		zap.S().Errorw("Error while marshalling http error",
			"error", err.Error())
	}
	return payload
}
