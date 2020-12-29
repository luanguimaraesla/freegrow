package session

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// generateJWT creates a new JWT token for a given interface
func generateJWT(uuid string, id interface{}, secret string, expiresAt time.Time) (string, error) {
	var err error

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uuid"] = uuid
	claims["id"] = id
	claims["exp"] = expiresAt.Unix()

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

// parseJWT checks if a JWT token is valid and parses its content
// into a *jwt.Token instance
func parseJWT(jwtToken, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// validateJWT returns an error if the token is not valid
func validateJWT(jwtToken, secret string) error {
	token, err := parseJWT(jwtToken, secret)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}

	return nil
}
