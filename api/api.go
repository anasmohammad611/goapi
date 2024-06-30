package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceRes struct {
	Balance int64
	Code    int
}

type Error struct {
	code    int
	message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		code:    code,
		message: message,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalServerErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
