package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	CreateUser(req RequestUser) (User, error)
	CheckExistEmail(req RequestUser) error
}

type services struct {
	repository Repository
}

func NewService(repository Repository) *services {
	return &services{repository}
}

func (s *services) CreateUser(req RequestUser) (User, error) {
	user := User{}
	user.Name = req.Name
	user.Email = req.Email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *services) CheckExistEmail(req RequestUser) error {
	email := req.Email

	if user := s.repository.FindEmail(email); user != nil {
		return errors.New("registered email")
	}

	return nil
}
