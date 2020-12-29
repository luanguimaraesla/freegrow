package session

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/luanguimaraesla/freegrow/internal/cache"
)

type Token struct {
	Uuid      string
	Id        interface{}
	Jwt       string
	ExpiresAt time.Time
}

// NewToken generates a new Token intance with an auto-generated Uuid
func NewToken(id interface{}, secret string, duration time.Duration) (*Token, error) {
	expiresAt := time.Now().Add(duration)
	tokenUuid := uuid.New().String()

	jwtToken, err := generateJWT(tokenUuid, id, secret, expiresAt)
	if err != nil {
		return nil, fmt.Errorf("unable to create a new session token: %v", err)
	}

	token := &Token{tokenUuid, id, jwtToken, expiresAt}

	return token, nil
}

// LoadToken receives a jwtToken and decodes its metadata to create
// a valid Token instance
func LoadToken(jwtToken, secret string) (*Token, error) {
	token, err := parseJWT(jwtToken, secret)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		tokenUuid, ok := claims["uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("unable to decode token's uuid: key not found")
		}

		id, ok := claims["id"]
		if !ok {
			return nil, fmt.Errorf("unable to decode token's id: key not found")
		}

		expString, ok := claims["exp"]
		if !ok {
			return nil, fmt.Errorf("unable to decode token's exp: key not found")
		}

		exp, err := strconv.ParseInt(fmt.Sprintf("%.f", expString), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("unable to parse JWT exp: %v", err)
		}

		t := &Token{
			Id:        id,
			Uuid:      tokenUuid,
			Jwt:       jwtToken,
			ExpiresAt: time.Unix(exp, 0),
		}

		return t, nil
	}

	return nil, fmt.Errorf("unable to decode JWT claims")
}

// Save persists token in the cache
func (t *Token) Save() error {
	now := time.Now()

	if err := cache.Setex(t.Uuid, fmt.Sprintf("%v", t.Id), t.ExpiresAt.Sub(now)); err != nil {
		return fmt.Errorf("unable to persist token into cache: %v", err)
	}

	return nil
}

// Delete removes token from cache
func (t *Token) Delete() error {
	if err := cache.Delete(t.Uuid); err != nil {
		return err
	}

	return nil
}

// Check checks if the token is still on the cache
func (t *Token) Check() (string, error) {
	id, err := cache.GetString(t.Uuid)
	if err != nil {
		return "", fmt.Errorf("unable to get session uuid in the cache: %v", err)
	}

	return id, nil
}
