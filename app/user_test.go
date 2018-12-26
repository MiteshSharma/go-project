package app

import (
	"testing"

	"github.com/MiteshSharma/project/model"
)

func TestCreateUser(t *testing.T) {
	at := SetupAppTest()
	defer at.CleanupAppTest()

	user := &model.User{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "random@gmail.com",
		Password:  "random",
	}
	t.Logf("Starting CreateUser test")
	userAuth, err := at.App.CreateUser(user)
	if err != nil {
		t.Logf("CreateUser test failed with err : %s", err.Message)
		t.FailNow()
	}
	if userAuth == nil {
		t.Logf("CreateUser test failed as no user auth created")
		t.FailNow()
	}
}
