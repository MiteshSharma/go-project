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

func TestUpdateUserDetail(t *testing.T) {
	t.Log("Starting update user detail test case")
	apiTest := GetApiTest()

	user := GetTestUser()
	userAuth := apiTest.CreateUserAuthFromTestAPI(t, apiTest.API, user)
	apiTest.CheckValidTestUser(t, user, userAuth.User)

	if userAuth.Token == "" {
		t.Errorf("handler returned wrong token: got %v",
			userAuth.Token)
	}

	expectedUser := userAuth.User
	userDetail := &model.UserDetail{
		UserID:    expectedUser.UserID,
		UtmSource: "Dummay",
	}
	t.Log(expectedUser)
	jsonUserDetail, _ := json.Marshal(userDetail)
	req, err := http.NewRequest("PUT", fmt.Sprintf("/api/v1/user/%d/userDetail", expectedUser.UserID), bytes.NewBuffer(jsonUserDetail))
	req.Header.Set("Authorization", userAuth.Token)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := apiTest.API.requestWithAuthHandler(apiTest.API.updateUserDetail)
	handler.ServeHTTP(res, req)

	CheckOkStatus(t, res.Code)

	t.Log(res.Body.String())
	receivedUserDetail := model.UserDetailFromString(res.Body.String())

	if receivedUserDetail.UserID != userDetail.UserID {
		t.Errorf("handler returned wrong name: got %d expected %d",
			receivedUserDetail.UserID, userDetail.UserID)
	}

	if receivedUserDetail.UtmSource != userDetail.UtmSource {
		t.Errorf("handler returned wrong utm source: got %s expected %s",
			receivedUserDetail.UtmSource, userDetail.UtmSource)
	}
}
