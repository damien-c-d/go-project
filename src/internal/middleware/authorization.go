package middleware

import (
	"errors"
	"golearn/src/api"
	"golearn/src/internal/tools"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var ErrorUnauthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(ErrorUnauthorized)
			api.RequestErrorHandler(w, ErrorUnauthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(ErrorUnauthorized)
			api.RequestErrorHandler(w, ErrorUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}
