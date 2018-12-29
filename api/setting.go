package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
)

func (a *API) InitSetting() {
	a.Router.Setting.Handle("", a.requestHandler(a.getSetting)).Methods("GET")
}

func (a *API) getSetting(rc *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	rc.SetAppResponse(rc.App.Setting.ToJson(), http.StatusOK)
}
