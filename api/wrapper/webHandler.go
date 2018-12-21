package wrapper

import (
	"net/http"

	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/model"
	uuid "github.com/satori/go.uuid"
)

type WebHandler struct {
	AppOption   *app.AppOption
	HandlerFunc func(*RequestContext, http.ResponseWriter, *http.Request)
}

func (wh *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rc := &RequestContext{}
	rc.App = app.NewApp(wh.AppOption)
	rc.RequestID = uuid.NewV4().String()
	rc.App.RequestID = rc.RequestID
	rc.Path = r.URL.Path

	w.Header().Set(model.HEADER_REQUEST_ID, rc.RequestID)

	wh.HandlerFunc(rc, w, r)

	if rc.Err != nil {
		w.Write([]byte(rc.Err.ToJson()))
		w.WriteHeader(rc.Err.Status)
		return
	}

	w.WriteHeader(200)
}
