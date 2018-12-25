package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
	"github.com/MiteshSharma/project/model"
)

func (a *API) InitUserDetail() {
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userDetail.ToJson()))
}
