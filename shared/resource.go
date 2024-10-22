package shared

import (
	"golang.org/x/crypto/bcrypt"
)

type Local struct {
	Salt int
}

func NewLocalConfig(salt int) *Local {
	return &Local{
		Salt: salt,
	}
}

func (local *Local) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), local.Salt)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
