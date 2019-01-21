package model

import (
	"encoding/json"
	"errors"
	"time"
)

// UserSession struct
type UserSession struct {
	UserSessionID int        `gorm:"primary_key" json:"userSessionId"`
	UserID        int        `gorm:"type:varchar(64)" json:"userId"`
	User          User       `gorm:"foreignkey:UserID" json:"-"`
	Token         string     `gorm:"type:varchar(1024)"`
	Roles         string     `gorm:"type:varchar(1024)"`
	CreatedAt     *time.Time `json:"-"`
	UpdatedAt     *time.Time `json:"-"`
	DeletedAt     *time.Time `json:"-"`
}

// Valid function is to check if policy object is valid
func (us *UserSession) Valid() error {
	if us.UserID == 0 {
		return errors.New("user id can not be 0")
	}
	return nil
}

func (us *UserSession) ToJson() string {
	json, _ := json.Marshal(us)
	return string(json)
}
