package httputils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
}

func GetRequestBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	return body, err
}

func WriteErrorResponse(w http.ResponseWriter, err Error) {
	errorTrace := map[string]string{
		"message": err.Message,
		"error":   err.Err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(errorTrace)
}
