package http_

import (
	"time"

	util "github.com/mirogon/go_util"
)

type Cookie struct {
	Name           string
	Value          string
	HttpOnly       bool
	SecureOnly     bool
	SameSite       string
	ExpirationDate *time.Time
}

func CreateCookie(name string, value string, httpOnly bool, secureOnly bool, sameSite string, expirationDate *time.Time) Cookie {
	return Cookie{Name: name, Value: value, HttpOnly: httpOnly, SecureOnly: secureOnly, SameSite: sameSite, ExpirationDate: expirationDate}
}

func (cookie Cookie) GetValueString() string {
	str := cookie.Name + "=" + cookie.Value + "; "

	if cookie.ExpirationDate != nil {
		str += "Expires=" + util.TimeInHttpFormat(*cookie.ExpirationDate) + "; "
	}

	if cookie.HttpOnly {
		str += "HttpOnly; "
	}

	if cookie.SecureOnly {
		str += "Secure; "
	}

	if cookie.SameSite != "" {
		str += "SameSite=" + cookie.SameSite + "; "
	}

	for str[len(str)-1] == ';' || str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}

	return str
}
