package utils

import (
	"github.com/nasim0x1/bifrost/configs"
	"golang.org/x/crypto/bcrypt"
)

func GenaratePasswordHash(password string) (string, error) {
	password = password + configs.Envs.Secret
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPasswordHash(hashedPassword, password string) bool {
	combined := password + configs.Envs.Secret
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(combined))
	return err == nil
}
