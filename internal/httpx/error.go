package httpx

import (
	"encoding/json"
	"net/http"
)

type Code string

const (
	INTERNAL_ERROR Code = "INTERNAL_ERROR"
	NOT_FOUND Code = "NOT_FOUND"
	BAD_REQUEST Code = "BAD_REQUEST"
	UNAUTHORIZED Code = "UNAUTHORIZED"
	FORBIDDEN Code = "FORBIDDEN"
	TOO_MANY_REQUESTS Code = "TOO_MANY_REQUESTS"
	INTERNAL_SERVER_ERROR Code = "INTERNAL_SERVER_ERROR"
	NOT_IMPLEMENTED Code = "NOT_IMPLEMENTED"
	BAD_GATEWAY Code = "BAD_GATEWAY"
	SERVICE_UNAVAILABLE Code = "SERVICE_UNAVAILABLE"
	GATEWAY_TIMEOUT Code = "GATEWAY_TIMEOUT"
	HTTP_VERSION_NOT_SUPPORTED Code = "HTTP_VERSION_NOT_SUPPORTED"
	VARIANT_ALSO_NEGOTIATES Code = "VARIANT_ALSO_NEGOTIATES"
	INSUFFICIENT_STORAGE Code = "INSUFFICIENT_STORAGE"
	LOOP_DETECTED Code = "LOOP_DETECTED"
	NOT_EXTENDED Code = "NOT_EXTENDED"
	NETWORK_AUTHENTICATION_REQUIRED Code = "NETWORK_AUTHENTICATION_REQUIRED"
)

type errorEnvalope struct {
	Error errorPayload `json:"error"`
}

type errorPayload struct {
	Message string `json:"message"`
	Code    Code   `json:"code"`
}

func Error(w http.ResponseWriter, status int, message string, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(errorEnvalope{Error: errorPayload{Message: message, Code: code}}) 
}
