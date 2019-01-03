package model

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

// UserDetail struct
type UserDetail struct {
	// the id for this user detail
	//
	// required: true
	// min: 1
	UserDetailID int `gorm:"primary_key" json:"userDetailId"`
	// the id for this user
	//
	// required: true
	// min: 1
	UserID int  `json:"userId"`
	User   User `gorm:"foreignkey:UserID" json:"-"`
	// source from where user came
	UtmSource string `gorm:"type:varchar(64)" json:"utmSource"`
	// campaign from where user came
	UtmCampaign string `gorm:"type:varchar(64)" json:"utmCampaign"`
	// medium from where user came
	UtmMedium string `gorm:"type:varchar(64)" json:"utmMedium"`
	// content from where user came
	UtmContent string     `gorm:"type:varchar(64)" json:"utmContent"`
	CreatedAt  *time.Time `json:"-"`
	UpdatedAt  *time.Time `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

// Valid function is to check if policy object is valid
func (ud *UserDetail) Valid() error {
	if ud.UserID == 0 {
		return errors.New("user id can not be 0")
	}
	return nil
}

func (ud *UserDetail) ToJson() string {
	json, _ := json.Marshal(ud)
	return string(json)
}

func UserDetailFromJson(data io.Reader) *UserDetail {
	var userDetail *UserDetail
	json.NewDecoder(data).Decode(&userDetail)
	return userDetail
}

func UserDetailFromString(data string) *UserDetail {
	var userDetail *UserDetail
	json.Unmarshal([]byte(data), &userDetail)
	return userDetail
}
