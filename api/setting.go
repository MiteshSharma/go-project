package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
)

func (a *API) InitSetting() {
	a.Router.Setting.Handle("", a.requestHandler(a.getSetting)).Methods("GET")
}

func (a *API) getSetting(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(c.App.Setting.ToJson()))
}
