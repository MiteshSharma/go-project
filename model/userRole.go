package model

import (
	"errors"
	"time"
)

// UserRole struct
type UserRole struct {
	UserRoleID int        `gorm:"primary_key"`
	UserID     int        `gorm:"type:varchar(64)"`
	User       User       `gorm:"foreignkey:UserID" json:"-"`
	Role       Role       `gorm:"type:varchar(64)"`
	CreatedAt  *time.Time `json:"-"`
	UpdatedAt  *time.Time `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

// Valid function is to check if policy object is valid
func (ur *UserRole) Valid() error {
	if ur.UserID == 0 {
		return errors.New("user id can not be 0")
	}
	if ur.Role == "" {
		return errors.New("role id can not be 0")
	}
	return nil
}
