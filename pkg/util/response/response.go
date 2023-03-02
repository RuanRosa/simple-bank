package response

import (
	"encoding/json"
	"net/http"
)

type Json struct{}

func (j *Json) Write(w http.ResponseWriter, payload interface{}, statusCode int) {
	respByte, err := json.Marshal(payload)
	if err != nil {
		j.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(respByte)
}

func (j *Json) WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	errorMsg := map[string]interface{}{
		"error:": err.Error(),
	}

	b, _ := json.Marshal(errorMsg)
	w.Write(b)
}
