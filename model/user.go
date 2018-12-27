package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// User struct
type User struct {
	UserID        int        `gorm:"primary_key" json:"userId"`
	FirstName     string     `gorm:"type:varchar(64)" json:"firstName"`
	LastName      string     `gorm:"type:varchar(64)" json:"lastName"`
	Email         string     `gorm:"type:varchar(100);unique_index" json:"email"`
	Password      string     `gorm:"type:varchar(256)" json:"-"`
	Salt          string     `gorm:"type:varchar(64)" json:"-"`
	ResetPassword string     `gorm:"type:varchar(32)" json:"-"`
	CreatedAt     *time.Time `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}

// Valid function is to check if policy object is valid
func (u *User) Valid() error {
	if u.Email == "" {
		return errors.New("user email can not be nil or empty")
	}
	return nil
}

func (u *User) ToJson() string {
	json, _ := json.Marshal(u)
	return string(json)
}

func UsersToJson(users []*User) string {
	json, _ := json.Marshal(users)
	return string(json)
}

func UserFromJson(data io.Reader) *User {
	var user *User
	json.NewDecoder(data).Decode(&user)
	return user
}

func UserFromString(data string) *User {
	var user *User
	json.Unmarshal([]byte(data), &user)
	return user
}

// UserAuth struct
type UserAuth struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

func (u *UserAuth) ToJson() string {
	json, _ := json.Marshal(u)
	return string(json)
}

func UserAuthFromString(data string) *UserAuth {
	var userAuth *UserAuth
	json.Unmarshal([]byte(data), &userAuth)
	return userAuth
}
