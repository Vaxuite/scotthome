package routing

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// SendPayload sends the payload back to the requesting client
func SendPayload(payload interface{}, w http.ResponseWriter) {
	payloadMarshal, err := json.Marshal(payload)
	if err != nil {
		RespondWithHTTPError("Error while marshalling payload", err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(payloadMarshal)
	if err != nil {
		zap.S().Errorw("Error writing response",
			"err", err.Error())
	}
}
