package model

import "encoding/json"

type AppError struct {
	Message   string
	Status    int
	RequestId string
}

func NewAppError(message string, status int) *AppError {
	err := &AppError{
		Message: message,
		Status:  status,
	}
	return err
}

func (u *AppError) ToJson() string {
	json, _ := json.Marshal(u)
	return string(json)
}
