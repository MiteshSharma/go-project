package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MiteshSharma/project/model"
)

func TestUserLogin(t *testing.T) {
	t.Log("Starting user login test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User
	expectedUser.Password = "random"
	jsonUser, _ := json.Marshal(expectedUser)
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/v1/user/%d/auth", expectedUser.UserID), bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestHandler(apiTest.API.userLogin)
	handler.ServeHTTP(res, req)

	CheckCreatedStatus(t, res.Code)

	t.Log(res.Body.String())
	receivedUserAuth := model.UserAuthFromString(res.Body.String())

	if receivedUserAuth.Token == "" {
		t.Errorf("handler returned wrong name: got %s", receivedUserAuth.Token)
	}
}

func TestUserLogout(t *testing.T) {
	t.Log("Starting user logout test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/v1/user/%d/auth", expectedUser.UserID), nil)
	req.Header.Set(model.AUTHENTICATION, userAuth.Token)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestWithAuthHandler(apiTest.API.userLogout)
	handler.ServeHTTP(res, req)

	CheckOkStatus(t, res.Code)

	t.Log(res.Body.String())

	if res.Body.String() != "{'response': 'OK'}" {
		t.Errorf("handler returned wrong response: got %s expected %s",
			res.Body.String(), "{'response': 'OK'}")
	}
}
