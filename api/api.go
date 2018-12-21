package api

import (
	"github.com/MiteshSharma/project/logger"
	"github.com/gorilla/mux"

	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/model"
)

type API struct {
	MainRouter *mux.Router
	AppOption  *app.AppOption
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	Router     *Router
}

func NewAPI(router *mux.Router, appOption *app.AppOption, config *model.Config, metrics metrics.Metrics, logger logger.Logger) *API {
	api := &API{
		MainRouter: router,
		AppOption:  appOption,
		Config:     config,
		Metrics:    metrics,
		Log:        logger,
		Router:     &Router{},
	}
	api.setupRoutes()
	return api
}
