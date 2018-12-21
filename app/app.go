package app

import (
	"github.com/MiteshSharma/project/bi"
	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/repository"
	"github.com/MiteshSharma/project/setting"
)

type App struct {
	Repository     repository.Repository
	Config         *model.Config
	Setting        *setting.Setting
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler *bi.BiEventHandler
	RequestID      string
}

func NewApp(appOption *AppOption) *App {
	app := &App{
		Repository:     appOption.Repository,
		Config:         appOption.Config,
		Setting:        appOption.Setting,
		Metrics:        appOption.Metrics,
		Log:            appOption.Log,
		BiEventHandler: appOption.BiEventHandler,
	}
	return app
}
