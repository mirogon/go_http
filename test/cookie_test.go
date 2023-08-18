package http_test

import (
	"testing"
	"time"

	http_ "github.com/mirogon/go_http"
)

func TestCookie(t *testing.T) {
	cookie := http_.CreateCookie("name", "value", true, false, nil)
	if cookie.GetValueString() != "name=value; HttpOnly" {
		t.Error("!" + cookie.GetValueString() + "!")
	}

	cookie = http_.CreateCookie("name", "value", true, true, nil)
	if cookie.GetValueString() != "name=value; HttpOnly; Secure" {
		t.Error("!" + cookie.GetValueString() + "!")
	}

	cookie = http_.CreateCookie("name", "value", false, false, nil)
	if cookie.GetValueString() != "name=value" {
		t.Error("!" + cookie.GetValueString() + "!")
	}

	loc, _ := time.LoadLocation("GMT")
	someTime := time.Date(2023, 1, 1, 0, 7, 9, 0, loc)
	cookie = http_.CreateCookie("name", "value", false, false, &someTime)
	if cookie.GetValueString() != "name=value; Expires=Sun, 1 Jan 2023 00:07:09 GMT" {
		t.Error("!" + cookie.GetValueString() + "!")
	}

	loc, _ = time.LoadLocation("GMT")
	someTime = time.Date(2023, 1, 1, 0, 7, 9, 0, loc)
	cookie = http_.CreateCookie("name", "value", true, true, &someTime)
	if cookie.GetValueString() != "name=value; Expires=Sun, 1 Jan 2023 00:07:09 GMT; HttpOnly; Secure" {
		t.Error("!" + cookie.GetValueString() + "!")
	}

}
