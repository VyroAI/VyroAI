package bcryptRepository

import (
	"github.com/vyroai/VyroAI/internal/domain/authentication/repo"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
}

func NewBcryptRepository() repo.BcryptRepo {
	return &Bcrypt{}
}

func (b *Bcrypt) CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (b *Bcrypt) GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
