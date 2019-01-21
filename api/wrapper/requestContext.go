package wrapper

import (
	"net/http"

	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/model"
)

type RequestContext struct {
	RequestID   string
	Path        string
	AppResponse *model.AppResponse
	App         *app.App
	Err         *model.AppError
}

func (rc *RequestContext) SetError(message string, statusCode int) {
	rc.Err = model.NewAppError(message, statusCode)
}

func (rc *RequestContext) SetPermissionError(permission *model.Permission) {
	rc.Err = model.NewAppError(permission.Description, http.StatusForbidden)
}

func (rc *RequestContext) SetAppResponse(response string, statusCode int) {
	rc.AppResponse = model.NewAppResponse(response, statusCode)
}

func (rc *RequestContext) IsSudoUser() bool {
	return rc.App.UserHasPermissionTo(model.PERMISSION_SUDO_USER.ID)
}

func (rc *RequestContext) GetSession(r *http.Request) (*model.UserSession, *model.AppError) {
	token, err := rc.GetToken(r)
	if err == nil {
		return rc.App.VerifyAndParseToken(token)
	}
	return nil, err
}

func (rc *RequestContext) GetToken(r *http.Request) (string, *model.AppError) {
	return r.Header.Get(model.AUTHENTICATION), nil
}
