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

type AppResponse struct {
	Response string
	Status   int
}

func NewAppResponse(response string, status int) *AppResponse {
	appResponse := &AppResponse{
		Response: response,
		Status:   status,
	}
	return appResponse
}

func (u *AppResponse) ToJson() string {
	json, _ := json.Marshal(u)
	return string(json)
}
