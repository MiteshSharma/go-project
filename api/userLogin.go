package api

import (
	"net/http"
	"strconv"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
	"github.com/gorilla/mux"
)

func (a *API) InitUserLogin() {
	a.Router.User.Handle("/{userId:[0-9]+}/auth", a.requestHandler(a.userLogin)).Methods("POST")
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

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(userAuth.ToJson()))
}

func (a *API) userLogout(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDVal := vars["userId"]
	if userIDVal == "" {
		rc.SetError("UserId received to update userID is invalid.", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDVal)
	if err != nil {
		rc.SetError("UserId value invalid, not a integer.", http.StatusBadRequest)
		return
	}
	rc.App.UserLogout(userID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{'response': 'OK'}"))
}
