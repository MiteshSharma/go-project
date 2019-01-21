package wrapper

import (
	"net/http"
	"time"

	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/model"
	uuid "github.com/satori/go.uuid"
)

type WebHandler struct {
	AppOption   *app.AppOption
	HandlerFunc func(*RequestContext, http.ResponseWriter, *http.Request)
	IsLoggedIn  bool
	IsSudoUser  bool
}

func (wh *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	rc := &RequestContext{}
	rc.App = app.NewApp(wh.AppOption)
	rc.RequestID = uuid.NewV4().String()
	rc.App.RequestID = rc.RequestID
	rc.Path = r.URL.Path

	w.Header().Set(model.HEADER_REQUEST_ID, rc.RequestID)

	if wh.IsLoggedIn {
		rc.App.UserSession, rc.Err = rc.GetSession(r)
	}

	if wh.IsSudoUser && !rc.IsSudoUser() {
		rc.SetPermissionError(model.PERMISSION_SUDO_USER)
	}

	if rc.Err == nil {
		wh.HandlerFunc(rc, w, r)
	}

	statusCode := http.StatusOK
	if rc.Err != nil {
		statusCode = rc.Err.Status
		rc.Err.RequestId = rc.RequestID
		w.Write([]byte(rc.Err.ToJson()))
	}
	if rc.AppResponse != nil {
		statusCode = rc.AppResponse.Status
	}
	w.WriteHeader(statusCode)
	if rc.AppResponse != nil {
		w.Write([]byte(rc.AppResponse.Response))
	}

	if rc.App.Metrics != nil {
		elapsedDuration := float64(time.Since(now).Nanoseconds()) / float64(time.Millisecond)
		rc.App.Metrics.RequestReceivedDetail(rc.Path, r.Method, statusCode, elapsedDuration)
	}
}
