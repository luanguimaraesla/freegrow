package session

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/luanguimaraesla/freegrow/internal/cache"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"go.uber.org/zap"
)

var (
	sessionDuration = 300 * time.Second
	cookieName      = "SESSION_TOKEN"
)

// CreateSession creates a new session and stores in the cache
func CreateSession(w http.ResponseWriter, i interface{}) error {

	// Create a new random session token
	sessionToken := uuid.New().String()

	if err := cache.Setex(sessionToken, i, sessionDuration); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    cookieName,
		Value:   sessionToken,
		Expires: time.Now().Add(sessionDuration),
	})

	return nil
}

// CheckSession is a decorator that checks session token and passes its
// value to the decorated router function
func CheckSession(fn func(string, http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		sessionToken := cookie.Value

		s, err := cache.GetString(sessionToken)
		if err != nil {
			log.L.Error("failed getting SESSION_TOKEN", zap.Error(err))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		fn(s, w, r)
	}
}
