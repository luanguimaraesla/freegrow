package session

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/luanguimaraesla/freegrow/internal/log"
	"go.uber.org/zap"
)

// CloseSession receives an JWT encoded token and finishes this session
func CloseSession(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := extractToken(r)
		if err != nil {
			log.L.Error("failed extracting request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := LoadToken(jwtToken, opts.accessSecret)
		if err != nil {
			log.L.Error("failed loading request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if err := token.Delete(); err != nil {
			log.L.Error("failed deleting request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fn(w, r)
	}
}

// CheckSession is a middleware that checks session token and passes its
// value to the decorated router function
func CheckSession(fn func(string, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken, err := extractToken(r)
		if err != nil {
			log.L.Error("failed extracting request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := LoadToken(jwtToken, opts.accessSecret)
		if err != nil {
			log.L.Error("failed loading request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		id, err := token.Check()
		if err != nil {
			log.L.Error("failed checking request token", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fn(id, w, r)
	}
}

// extractToken gets the token from a user's request header
func extractToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearerToken, " ")
	if len(strArr) != 2 {
		return "", fmt.Errorf("unable to decode bearerToken: invalid format")
	}

	return strArr[1], nil
}
