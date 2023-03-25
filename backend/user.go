package main

import (
	"encoding/json"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	FirstName string
	LastName  string
	Password  string `json:"-"`
	Role      uint   `json:"-"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type userAlias User // Define an alias of the struct to avoid infinite recursion
	aux := &struct {
		*userAlias
		Password string `json:"password"`
	}{
		userAlias: (*userAlias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Password != "" {
		u.Password = aux.Password
	}
	return nil
}