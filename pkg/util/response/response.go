package response

import (
	"encoding/json"
	"net/http"
)

type Json struct{}

func (j *Json) Write(w http.ResponseWriter, payload interface{}, statusCode int) {
	body := []byte{}

	if payload != nil {
		resp, err := json.Marshal(payload)
		if err != nil {
			j.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		body = resp
	}

	w.WriteHeader(statusCode)
	w.Write(body)
}

func (j *Json) WriteError(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	body := []byte{}
	if err != nil {
		errorMsg := map[string]interface{}{
			"error:": err.Error(),
		}

		body, _ = json.Marshal(errorMsg)
	}

	w.Write(body)
}
