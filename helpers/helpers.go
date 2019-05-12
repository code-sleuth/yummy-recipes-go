package helpers

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse runc
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	JSONResponse(w, statusCode, map[string]string{"error": message})
}

// JSONResponse func
func JSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
