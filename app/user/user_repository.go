package user

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	InsertUser(user User) (User, error)
	FindEmail(email string) *User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	err := r.db.Create(&user)
	if err != nil {
		return user, err.Error
	}

	return user, nil
}

func (r *repository) FindEmail(email string) *User {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil
	}

	return &user
}
