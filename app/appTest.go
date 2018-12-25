package app

import (
	"fmt"
	"time"

	"github.com/MiteshSharma/project/bus"

	"github.com/MiteshSharma/project/bi"

	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/setting"
)

type AppTest struct {
	App *App
}

func Setup() *AppTest {
	now := time.Now()
	startTime := fmt.Sprintf("%d", now.Unix())
	settingData := setting.NewSetting("1", "1", "NA", "master", startTime)
	config := setting.GetConfig()

	appOption := &AppOption{
		Config:         config,
		Setting:        settingData,
		Log:            logger.NewTestLogger(config),
		Metrics:        metrics.NewTestMetrics(),
		Repository:     nil,
		BiEventHandler: bi.NewBiTestEventHandler(),
		Bus:            bus.NewTestBus(),
	}
	app := NewApp(appOption)

	appTest := &AppTest{
		App: app,
	}
	return appTest
}

func (at *AppTest) Cleanup() {
}
