package model

import (
	"errors"
	"time"
)

// UserDetail struct
type UserDetail struct {
	UserDetailID int        `gorm:"primary_key" json:"userDetailId"`
	UserID       int        `json:"userId"`
	User         User       `gorm:"foreignkey:UserID"`
	UtmSource    string     `gorm:"type:varchar(64)" json:"utmSource"`
	UtmCampaign  string     `gorm:"type:varchar(64)" json:"utmCampaign"`
	UtmMedium    string     `gorm:"type:varchar(64)" json:"utmMedium"`
	UtmContent   string     `gorm:"type:varchar(64)" json:"utmContent"`
	CreatedAt    *time.Time `json:"-"`
	UpdatedAt    *time.Time `json:"-"`
	DeletedAt    *time.Time `json:"-"`
}

// Valid function is to check if policy object is valid
func (ud *UserDetail) Valid() error {
	if ud.UserID == 0 {
		return errors.New("user id can not be 0")
	}
	return nil
}
