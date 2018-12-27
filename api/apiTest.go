package api

import (
	"fmt"
	"time"

	"github.com/MiteshSharma/project/repository/docker"

	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/bi"
	"github.com/MiteshSharma/project/bus"
	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/repository"
	"github.com/MiteshSharma/project/setting"
	"github.com/gorilla/mux"
)

type APITest struct {
	API         *API
	MySQLDocker *docker.MysqlDocker
}

var apiTest *APITest

func SetupApiTest() *APITest {
	now := time.Now()
	startTime := fmt.Sprintf("%d", now.Unix())
	settingData := setting.NewSetting("1", "1", "NA", "master", startTime)
	config := setting.GetConfigFromFile("test")
	logger := logger.NewTestLogger(config)
	metrics := metrics.NewTestMetrics()
	biEventHandler := bi.NewBiTestEventHandler()
	bus := bus.NewTestBus()
	router := mux.NewRouter()

	mysqlDocker := &docker.MysqlDocker{
		ContainerName: "mysql-api-container",
	}
	fmt.Println("Starting docker mysql container")
	mysqlDocker.StartMysqlDocker()
	fmt.Println("Started docker mysql container")
	fmt.Println("waiting for 10 sec before start")
	// Wait for docker mysql server to start
	time.Sleep(10 * time.Second)
	fmt.Println("waiting complete")

	repository := repository.NewPersistentRepository(logger, config, metrics)

	appOption := &app.AppOption{
		Config:         config,
		Setting:        settingData,
		Log:            logger,
		Metrics:        metrics,
		Repository:     repository,
		BiEventHandler: biEventHandler,
		Bus:            bus,
	}

	api := NewAPI(router, appOption, config, metrics, logger)
	apiTest = &APITest{
		API:         api,
		MySQLDocker: mysqlDocker,
	}
	return apiTest
}

func GetApiTest() *APITest {
	return apiTest
}

func (at *APITest) CleanUpApiTest() {
	at.MySQLDocker.Stop()
}
