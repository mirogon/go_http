package http_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	http_ "github.com/mirogon/go_http"
	mock_http_ "github.com/mirogon/go_http/mocks"
)

type LoginRequest struct {
	Email      string
	Password   string
	LoginToken string
}

type LoginResponse struct {
	ErrorCode  int    `json:"errorCode"`
	Username   string `json:"username"`
	LoginToken string `json:"loginToken"`
	Permission int    `json:"permission"`
	UserSince  string `json:"userSince"`
	TokenLeft  int    `json:"tokenLeft"`
}

func TestSendLoginResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	responseWriterMock := mock_http_.NewMockHttpResponseWriter(ctrl)
	loginResponse := LoginResponse{}
	jsonData, _ := json.Marshal(loginResponse)
	responseWriterMock.EXPECT().SetBody(jsonData)
	responseWriterMock.EXPECT().Send()

	http_.SendHttpResponse(responseWriterMock, loginResponse)
}

func TestHandleOptionsRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	responseWriterMock := mock_http_.NewMockHttpResponseWriter(ctrl)
	responseWriterMock.EXPECT().AddHeader("Access-Control-Allow-Origin", gomock.Any())
	responseWriterMock.EXPECT().AddHeader("Access-Control-Allow-Headers", "Content-Type, Authorization")
	responseWriterMock.EXPECT().AddHeader("Access-Control-Allow-Credentials", "true")
	responseWriterMock.EXPECT().DeleteHeader("Content-Type")
	responseWriterMock.EXPECT().DeleteHeader("Content-Length")
	responseWriterMock.EXPECT().Send()

	http_.HandlePreflightOptionsRequest(responseWriterMock, "")
}

func TestParseRequestIntoRequestType(t *testing.T) {
	reader := strings.NewReader("{\"email\": \"test@example.com\", \"password\": \"123456\", \"logintoken\": \"123\"}")
	request, _ := http.NewRequest("any", "/", reader)
	result, _ := http_.ParseRequestIntoRequestType[LoginRequest](request)
	if result.Email != "test@example.com" {
		t.Error()
	}
	if result.Password != "123456" {
		t.Error()
	}
	if result.LoginToken != "123" {
		t.Error()
	}
}
