package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUserLogin() {
	// swagger:operation POST /user/{userId}/auth userLogin users
	// ---
	// summary: login user
	// description: Send email with password to authenticate.
	// parameters:
	// - name: userId
	//   in: path
	//   description: unique identifier of user
	//   type: int
	//   required: true
	// - name: email
	//   in: body
	//   description: email of user
	//   type: email
	//   required: true
	//   example: user@goproject.com
	// - name: password
	//   in: body
	//   description: password of user
	//   type: string
	//   required: true
	// Responses:
	//  201:
	//   description: success return user
	//   schema:
	//     "$ref": '#/definitions/UserAuth'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("/{userId:[0-9]+}/auth", a.requestHandler(a.userLogin)).Methods("POST")
	// swagger:operation DELETE /user/{userId}/auth userLogin users
	// ---
	// summary: Logout user
	// description: Logout user
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
	// Responses:
	//  200:
	//   description: deleted successfully return response ok
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("/{userId:[0-9]+}/auth", a.requestWithAuthHandler(a.userLogout)).Methods("DELETE")
}

func (a *API) userLogin(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if err := user.Valid(); err != nil {
		rc.SetError("User object received is not valid.", http.StatusBadRequest)
		return
	}

	userAuth, err := rc.App.UserLogin(user)
	if err != nil {
		rc.SetError("User login error.", http.StatusInternalServerError)
		return
	}

	rc.SetAppResponse(userAuth.ToJson(), http.StatusCreated)
}

func (a *API) userLogout(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.App.UserSession.UserID
	if userID == 0 {
		rc.SetError("UserId received is invalid.", http.StatusBadRequest)
		return
	}
	rc.App.UserLogout(userID)

	rc.SetAppResponse("{'response': 'OK'}", http.StatusOK)
}
