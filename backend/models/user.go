package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Email      string `gorm:"size: 100;not null;unique" json:"user"`
	FirstName  string `gorm:"size: 50;not null" json:"firstName"`
	LastName   string `gorm:"size: 50;not null" json:"lastName"`
	Password   string `json:"-"`
	Role       uint   `json:"-"`
}

func (user *User) HashPassword() error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) Exists() bool {
	db := GetDB()
	var count int64
	db.Model(&User{}).Where("email = ?", user.Email).Count(&count)
	return count > 0
}

func (user *User) Save() {
	GetDB().Create(&user)
}


