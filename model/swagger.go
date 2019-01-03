package model

// An UserAuthResponse response model
//
// This is used for returning a response with a single order as body
//
// swagger:response UserAuthResponse
type UserAuthResponse struct {
	// in: body
	Payload *UserAuth `json:"userAuth"`
}

// An UserDetailResponse response model
//
// This is used for returning a response with a single user detail as body
//
// swagger:response UserDetailResponse
type UserDetailResponse struct {
	// in: body
	Payload *UserDetail `json:"userDetail"`
}

// An AppErrorResponse response model
//
// This is used for returning a response with a single order as body
//
// swagger:response AppErrorResponse
type AppErrorResponse struct {
	// in: body
	Payload *AppError `json:"appError"`
}
