package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/luanguimaraesla/freegrow/internal/cache"
)

var (
	sessionDuration = 300 * time.Second
	cookieName      = "session_token"
)

// CreateSession creates a new session and stores in the cache
func CreateSession(w http.ResponseWriter, i interface{}) error {
	c, err := cache.Connect()
	if err != nil {
		return err
	}

	defer c.Close()

	// Create a new random session token
	sessionToken := uuid.New().String()

	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	duration := fmt.Sprintf("%.0f", sessionDuration.Seconds())
	if _, err := c.Do("SETEX", sessionToken, duration, fmt.Sprintf("%", i)); err != nil {
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

		c, err := cache.Connect()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer c.Close()

		response, err := c.Do("GET", sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if response == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		s, ok := response.(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fn(s, w, r)
	}
}
