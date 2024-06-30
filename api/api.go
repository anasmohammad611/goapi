package api

import (
	"encoding/json"
	"fmt"
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
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalServerErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal server error", http.StatusInternalServerError)
	}
)
