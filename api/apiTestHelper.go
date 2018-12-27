package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MiteshSharma/project/model"
	uuid "github.com/satori/go.uuid"
)

func CheckCreatedStatus(t *testing.T, statusCode int) {
	t.Helper()

	if statusCode != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			statusCode, http.StatusCreated)
	}
}

func CheckOkStatus(t *testing.T, statusCode int) {
	t.Helper()

	if statusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			statusCode, http.StatusOK)
	}
}

func GetTestUniqueEmail() string {
	uid := uuid.NewV4()
	return fmt.Sprintf("%s@testemail.com", uid)
}

func GetTestUser() *model.User {
	user := &model.User{
		FirstName: "Test",
		LastName:  "Test",
		Email:     GetTestUniqueEmail(),
		Password:  "random",
	}
	return user
}

func CheckValidTestUser(t *testing.T, expectedUser *model.User, receivedUser *model.User) {
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

func CreateUserAuthFromTestAPI(t *testing.T, api *API, user *model.User) *model.UserAuth {
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
