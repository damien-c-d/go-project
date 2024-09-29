package api

import (
	"encoding/json"
	"net/http"
)

type PointBalanceParams struct {
	Username string
}

type PointBalanceResponse struct {
	Code    int
	Balance int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, http.StatusNotFound, "Invalid Request")
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "Internal Server Error")
	}
)
