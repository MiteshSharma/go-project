package wrapper

import (
	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/model"
)

type RequestContext struct {
	RequestID string
	Path      string
	App       *app.App
	Err       *model.AppError
}

func (rc *RequestContext) SetError(message string, statusCode int) {
	rc.Err = model.NewAppError(message, statusCode)
}
