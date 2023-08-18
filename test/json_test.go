package http_test

import (
	"encoding/json"
	"testing"
	"time"

	util "github.com/mirogon/go_util"
)

func TestJsonParse(t *testing.T) {
	timeString := util.TimeInHttpFormat(time.Now())
	loginResponse := LoginResponse{Username: "bob", Permission: 0, UserSince: timeString}

	_, err := json.Marshal(loginResponse)
	if err != nil {
		t.Error(err)
	}

}
