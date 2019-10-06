package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	PhoneNumber string
	Password    string
	NickName    string
	Bio         string
	Status      string
	Role        string
	Avatar      string
}

const (
	PasswordCosts        = 12
	Active        string = "active"
	Inactive      string = "inactive"
	Suspend       string = "suspend"
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCosts)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
