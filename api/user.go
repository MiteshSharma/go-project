package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUser() {
	// swagger:operation POST /user user users
	// ---
	// summary: Create new user
	// description: Send email with other user details to create new user. Email must be unique.
	// parameters:
	// - name: body
	//   in: body
	//   schema:
	//      $ref: '#/definitions/User'
	// - name: password
	//   in: body
	//   description: password of user
	//   type: string
	//   required: true
	// Responses:
	//  201:
	//   description: success return user auth containing user with auth token
	//   schema:
	//     "$ref": '#/definitions/UserAuth'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("", a.requestHandler(a.createUser)).Methods("POST")
	// swagger:operation PUT /user/{userId} user users
	// ---
	// summary: Update user
	// description: Send user body with userId with other user data which needs to be updated.
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
	//   description: updated user body, must have userId non zero
	//   schema:
	//      $ref: '#/definitions/User'
	// Responses:
	//  200:
	//   description: success return updated user
	//   schema:
	//     "$ref": '#/definitions/User'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.updateUser)).Methods("PUT")
	// swagger:operation DELETE /user/{userId} user users
	// ---
	// summary: Delete user
	// description: Delete user entry from backend.
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
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.deleteUser)).Methods("DELETE")
	// swagger:operation GET /user/{userId} user users
	// ---
	// summary: Get user object
	// description: Get user object based on unique user identifier.
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
	//   description: success return user data
	//   schema:
	//     "$ref": '#/definitions/User'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.getUser)).Methods("GET")
	// swagger:operation GET /user user users
	// ---
	// summary: Get all user objects
	// description: Get all user objects created
	// Security:
	// - AuthKey: []
	// parameters:
	// - name: Authorization
	//   in: header
	//   description: JTW token used to validate user
	//   type: string
	//   required: true
	// Responses:
	//  200:
	//   description: success return user data
	//   schema:
	//	   type: array
	//	   items:
	//       "$ref": '#/definitions/User'
	//  default:
	//   description: unexpected error
	//   schema:
	//     "$ref": "#/definitions/AppError"
	a.Router.User.Handle("", a.requestWithAuthHandler(a.getAllUser)).Methods("GET")
}

// CreateHandler func is used to create user
func (a *API) createUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if err := user.Valid(); err != nil {
		rc.SetError("User object received is not valid.", http.StatusBadRequest)
		return
	}
	var userAuth *model.UserAuth
	var err *model.AppError
	if userAuth, err = rc.App.CreateUser(user); err != nil {
		rc.SetError("User object creation failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(userAuth.ToJson(), http.StatusCreated)
}

// UpdateHandler func is used to create user
func (a *API) updateUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		rc.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if user.UserID == 0 {
		rc.SetError("UserId received to update userID is 0.", http.StatusBadRequest)
		return
	}
	var err *model.AppError
	if user, err = rc.App.UpdateUser(user); err != nil {
		rc.SetError("User object update failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse(user.ToJson(), http.StatusOK)
}

// GetHandler func is used to get user or users
func (a *API) getUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.App.UserSession.UserID
	var user *model.User
	var appErr *model.AppError
	if user, appErr = rc.App.GetUser(userID); appErr != nil {
		rc.SetError("User object get failed.", http.StatusInternalServerError)
		return
	}

	rc.SetAppResponse(user.ToJson(), http.StatusOK)
}

// DeleteHandler func is to delete user
func (a *API) deleteUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.App.UserSession.UserID
	if _, err := rc.App.DeleteUser(userID); err != nil {
		rc.SetError("User object get failed.", http.StatusInternalServerError)
		return
	}
	rc.SetAppResponse("{'response': 'OK'}", http.StatusOK)
}

// GetHandler func is used to get user or users
func (a *API) getAllUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	var users []*model.User
	var err *model.AppError
	if users, err = rc.App.GetAllUser(); err != nil {
		rc.SetError("All users object get failed.", http.StatusInternalServerError)
		return
	}

	rc.SetAppResponse(model.UsersToJson(users), http.StatusOK)
}
