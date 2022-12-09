package user

import (
	"context"
	"fmt"
	"strings"
	"unicode"
)

func (s *Service) Save(ctx context.Context, data User) (int, error) {
	err := validateUser(s, ctx, data)
	if err != nil {
		return 0, err
	}

	var id int
	row, err := s.db.NamedQuery("INSERT INTO users(email, password, type) VALUES (:email, :password, :type) RETURNING id", data)
	if err != nil {
		return 0, err
	}
	if row.Next() {
		row.Scan(&id)
	}
	return id, nil
}

func validateUser(s *Service, ctx context.Context, data User) error {
	emailValidator := &EmailValidator{s, ctx}
	err := emailValidator.validate(data.Email)
	if err != nil {
		return err
	}

	passwordValidator := &PasswordValidator{}
	err = passwordValidator.validate(data.Password)
	if err != nil {
		return err
	}

	return nil
}

type Validator interface {
	validate(string) error
}

type EmailValidator struct {
	service *Service
	ctx     context.Context
}

func (validator *EmailValidator) validate(email string) error {
	isValidEmail := strings.Contains(email, "@")
	if !isValidEmail {
		return fmt.Errorf("email not valid")
	}

	var validationUser User
	_ = validator.service.db.GetContext(validator.ctx, &validationUser, "SELECT * FROM users WHERE email = $1", email)

	if validationUser.Email == email {
		return fmt.Errorf("User already exists")
	}
	return nil
}

type PasswordValidator struct {
}

func (validator *PasswordValidator) validate(password string) error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 7 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial {
		return nil
	} else {
		return fmt.Errorf("the password must have at least one upper and lower case char, one number, and 1 symbol")
	}
}
