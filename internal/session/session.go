package session

import (
	"fmt"
	"time"
)

var (
	opts *sessionOptions
)

// sessionOptions represents configurable values for
// access session management
type sessionOptions struct {
	cookieName           string
	accessSecret         string
	refreshSecret        string
	refreshTokenDuration time.Duration
	accessTokenDuration  time.Duration
}

type Session struct {
	AccessToken  *Token
	RefreshToken *Token
}

// Init defines JWT access secret key and session duration
func Init(accessSecret, refreshSecret string, accessTokenDuration, refreshTokenDuration time.Duration) {
	opts = &sessionOptions{
		cookieName:           "SESSION_TOKEN",
		accessSecret:         accessSecret,
		refreshSecret:        refreshSecret,
		accessTokenDuration:  accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

// New creates a new session fon a given id
func New(id interface{}) (*Session, error) {
	accessToken, err := NewToken(id, opts.accessSecret, opts.accessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("unable to create session access token: %v", err)
	}

	refreshToken, err := NewToken(id, opts.refreshSecret, opts.refreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("unable to create session refresh token: %v", err)
	}

	session := &Session{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return session, nil
}

// CreateSession creates a new session and returns the JWT access and refresh token
func CreateSession(id interface{}) (string, string, error) {
	session, err := New(id)
	if err != nil {
		return "", "", err
	}

	if err := session.Save(); err != nil {
		return "", "", err
	}

	return session.AccessToken.Jwt, session.RefreshToken.Jwt, nil
}

// Save persists both accessToken and refreshToken in the cache
func (s *Session) Save() error {
	if err := s.AccessToken.Save(); err != nil {
		return err
	}

	if err := s.RefreshToken.Save(); err != nil {
		return err
	}

	return nil
}
