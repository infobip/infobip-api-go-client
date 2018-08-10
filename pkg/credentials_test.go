package infobip

import "testing"

func TestBasicCredentials(t *testing.T) {
	b := NewBasicCredentials("username", "password")
	a := b.Authorization()
	e := "Basic dXNlcm5hbWU6cGFzc3dvcmQ="
	if a != e {
		t.Errorf("Expected basic authorization to be `%s`, but was `%s`", e, a)
	}
}

func TestApiKeyCredentials(t *testing.T) {
	k := NewApiKeyCredentials("api-key")
	a := k.Authorization()
	e := "App api-key"
	if a != e {
		t.Errorf("Expected api key authorization to be `%s`, but was `%s`", e, a)
	}
}

func TestTokenCredentilas(t *testing.T) {
	tk := NewIBSSOCredentials("ibsso-token")
	a := tk.Authorization()
	e := "IBSSO ibsso-token"
	if a != e {
		t.Errorf("Expected token authorization to be `%s`, but was `%s`", e, a)
	}
}
