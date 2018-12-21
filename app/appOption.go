package app

import (
	"github.com/MiteshSharma/project/bi"
	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/repository"
	"github.com/MiteshSharma/project/setting"
)

type AppOption struct {
	Repository     repository.Repository
	Config         *model.Config
	Setting        *setting.Setting
	Metrics        metrics.Metrics
	Log            logger.Logger
	BiEventHandler *bi.BiEventHandler
}
