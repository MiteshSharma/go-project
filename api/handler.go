package api

import (
	"net/http"

	"github.com/MiteshSharma/project/api/wrapper"
)

func (api *API) requestHandler(handler func(c *wrapper.RequestContext, w http.ResponseWriter, r *http.Request)) http.Handler {
	return &wrapper.WebHandler{
		AppOption:   api.AppOption,
		HandlerFunc: handler,
	}
}
