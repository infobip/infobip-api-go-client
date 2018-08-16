package sms

//This is a generated file and is not intended for modification!

import (
	"net/http"
	"net/http/httptest"
	"testing"

	infobip "github.com/infobip/infobip-api-go-client/pkg"
)

var (
	//reference imported packages just in case auto-generated code doesn't
	_ = http.StatusOK
	_ = httptest.DefaultRemoteAddr
	_ = testing.RunTests
	_ = infobip.NewApiKeyCredentials
)

func TestGetSentSMSDeliveryReports(t *testing.T) {
	ts := givenServer()
	defer ts.Close()
	c := givenClient(ts)

	_, err := c.GetSentSMSDeliveryReports(GetSentSMSDeliveryReportsQuery{})

	if err != nil {
		t.Errorf("Expected to GetSentSMSDeliveryReports, but was error %+v", err)
	}
}

func TestSendMultipleTextualSMSAdvanced(t *testing.T) {
	ts := givenServer()
	defer ts.Close()
	c := givenClient(ts)

	_, err := c.SendMultipleTextualSMSAdvanced(SMSAdvancedTextualRequest{})

	if err != nil {
		t.Errorf("Expected to SendMultipleTextualSMSAdvanced, but was error %+v", err)
	}
}

func givenServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}

func givenClient(ts *httptest.Server) Client {
	return Client{
		BaseURL: ts.URL,
		Authorizer: infobip.AuthorizerFunc(func() string {
			return "Auth"
		}),
	}
}
