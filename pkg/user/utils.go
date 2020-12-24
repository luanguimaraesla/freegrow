package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	if err != nil {
		return string(hashed), fmt.Errorf("failed to encrypt string: %v", err)
	}

	return string(hashed), nil
}
