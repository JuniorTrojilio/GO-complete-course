package validations

import (
	"api/src/models"
	"api/src/presentation"
	"errors"
)

// Validate receves a user and returns an error
func Validate(user *models.User) error {
	if user.Name == "" {
		return errors.New("O Nome é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O Nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}

	if user.Password == "" {
		return errors.New("A Senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// Prepare user to be call
func Prepare(user *models.User) error {
	if err := Validate(user); err != nil {
		return err
	}

	presentation.RemoveBlankSpace(user)

	return nil
}
