package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUser() {
	a.Router.User.Handle("", a.requestHandler(a.createUser)).Methods("POST")
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.updateUser)).Methods("PUT")
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.deleteUser)).Methods("DELETE")
	a.Router.User.Handle("/{userId:[0-9]+}", a.requestWithAuthHandler(a.getUser)).Methods("GET")
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
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(userAuth.ToJson()))
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.ToJson()))
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

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(user.ToJson()))
}

// DeleteHandler func is to delete user
func (a *API) deleteUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	userID := rc.App.UserSession.UserID
	if _, err := rc.App.DeleteUser(userID); err != nil {
		rc.SetError("User object get failed.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{'response': 'OK'}"))
}

// GetHandler func is used to get user or users
func (a *API) getAllUser(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	var users []*model.User
	var err *model.AppError
	if users, err = rc.App.GetAllUser(); err != nil {
		rc.SetError("All users object get failed.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(model.UsersToJson(users)))
}
