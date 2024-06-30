package main

import (
	"fmt"
	"github.com/anasmohammad611/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	log.SetReportCaller(true)
	var r = chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
