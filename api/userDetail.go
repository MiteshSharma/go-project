package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUserDetail() {
	// swagger:operation PUT /user/{userId}/userDetail userDetail users
	// ---
	// summary: Update user details
	// description: Send user detail identifier with userId which needs to be updated.
	// Security:
	// - AuthKey: []
	// parameters:
	// - name: Authorization
	//   in: header
	//   description: JTW token used to validate user
	//   type: string
	//   required: true
	// - name: userId
	//   in: path
	//   description: unique identifier of user
	//   type: int
	//   required: true
	// - name: body
	//   in: body
	//   description: updated user details, must have userId and userDetailId non zero
	//   schema:
	//      $ref: '#/definitions/UserDetail'
	// Responses:
	//  200:
	//   description: success return updated user
	//   schema:
	//     "$ref": '#/definitions/UserDetail'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("/{userId:[0-9]+}/userDetail", a.requestWithAuthHandler(a.updateUserDetail)).Methods("PUT")
}

func (a *API) updateUserDetail(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userDetail := model.UserDetailFromJson(r.Body)
	if userDetail == nil {
		rc.SetError("Body received for user detail is invalid.", http.StatusBadRequest)
		return
	}
	if userDetail.UserID == 0 {
		rc.SetError("UserId received to update user detail is 0.", http.StatusBadRequest)
		return
	}
	var err *model.AppError
	if userDetail, err = rc.App.UpdateUserDetail(userDetail); err != nil {
		rc.SetError("User object update failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(userDetail.ToJson(), http.StatusOK)
}
