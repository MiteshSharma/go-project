package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MiteshSharma/project/model"

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
	time.Sleep(10 * time.Second)
}

func (at *APITest) CheckValidTestUser(t *testing.T, expectedUser *model.User, receivedUser *model.User) {
	t.Helper()

	if expectedUser.Email != receivedUser.Email {
		t.Errorf("handler returned wrong email: got %v want %v",
			receivedUser.Email, expectedUser.Email)
	}

	if expectedUser.FirstName != receivedUser.FirstName {
		t.Errorf("handler returned wrong first name: got %v want %v",
			receivedUser.FirstName, expectedUser.FirstName)
	}

	if expectedUser.LastName != receivedUser.LastName {
		t.Errorf("handler returned wrong last name: got %v want %v",
			receivedUser.LastName, expectedUser.LastName)
	}
}

func (at *APITest) CreateUserAuthFromTestAPI(t *testing.T, api *API, user *model.User) *model.UserAuth {
	t.Helper()

	jsonUser, _ := json.Marshal(user)
	req, err := http.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestHandler(apiTest.API.createUser)
	handler.ServeHTTP(res, req)

	CheckCreatedStatus(t, res.Code)

	t.Logf("Create user response %s", res.Body.String())

	return model.UserAuthFromString(res.Body.String())
}
