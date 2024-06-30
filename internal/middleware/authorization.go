package middleware

import (
	"errors"
	"github.com/anasmohammad611/goapi/api"
	"github.com/anasmohammad611/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var unauthorizedError = errors.New("invalid token")
var paramNotFoundError = errors.New("param not found")
var userNotFoundError = errors.New("user not found")

func Authorization(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if len(username) == 0 {
			log.Error(paramNotFoundError)
			api.RequestErrorHandler(w, paramNotFoundError)
			return
		}

		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalServerErrorHandler(w)
			return
		}

		loginDetails, flag := database.GetUserLoginDetails(username)

		if flag == false {
			log.Error(userNotFoundError)
			api.RequestErrorHandler(w, userNotFoundError)
			return
		}

		if token != loginDetails.AuthToken || len(token) == 0 {
			log.Error(unauthorizedError)
			api.RequestErrorHandler(w, unauthorizedError)
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}
