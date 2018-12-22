package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUser() {
	a.Router.User.Handle("", a.requestHandler(a.create)).Methods("POST")
}

func (a *API) create(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	user := model.UserFromJson(r.Body)
	if user == nil {
		c.SetError("Body received for user creation is invalid.", http.StatusBadRequest)
		return
	}
	if err := user.Valid(); err != nil {
		a.Log.Warn("User object received is not valid.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := c.App.CreateUser(user); err != nil {
		a.Log.Warn("User object creation failed.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user.UserID == 0 {
		a.Log.Warn("User object creation failed as id is 0.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(user.ToJson()))
}
