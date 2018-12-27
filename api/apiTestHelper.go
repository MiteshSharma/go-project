package api

import (
	"fmt"
	"net/http"
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
