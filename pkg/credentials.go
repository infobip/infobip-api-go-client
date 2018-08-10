package infobip

import (
	"encoding/base64"
)

type Authorizer interface {
	Authorization() string
}

type AuthorizerFunc func() string

func (f AuthorizerFunc) Authorization() string {
	return f()
}

func NewBasicCredentials(username string, password string) Authorizer {
	s := username + ":" + password
	e := base64.StdEncoding.EncodeToString([]byte(s))
	a := "Basic " + e
	return AuthorizerFunc(func() string {
		return a
	})
}

func NewApiKeyCredentials(apiKey string) Authorizer {
	a := "App " + apiKey
	return AuthorizerFunc(func() string {
		return a
	})
}

func NewIBSSOCredentials(token string) Authorizer {
	a := "IBSSO " + token
	return AuthorizerFunc(func() string {
		return a
	})
}
