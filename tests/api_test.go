package infobip

import (
	"net/http"
	"os"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	ApiKeyIdentifier string = "Authorization"
	ApiKey           string = "App e64bb3cb13ff67fff448f0309d305338-f11b8706-5af1-44cf-b904-ec587e0f7129"
	Host             string = "localhost:8080"
)

var configuration infobip.Configuration
var infobipClient api.APIClient

var givenRequestHeaders = map[string]string{
	"Content-Type": "application/json;charset=UTF-8",
	"server":       "SMS API",
	"x-request-id": "1608758729810312842",
}

func TestMain(m *testing.M) {
	SetUpConfiguration()
	SetUpClient()

	os.Exit(m.Run())
}

func SetUpConfiguration() {
	configuration = *infobip.NewConfiguration()
	configuration.Scheme = "https"
	configuration.Host = Host
	configuration.DefaultHeader[ApiKeyIdentifier] = ApiKey
}

func SetUpClient() {
	infobipClient = *api.NewAPIClient(&configuration)
}

func SetUpSuccessRequest(requestType string, endpoint string, expectedResponse string, statusCode int) {
	httpmock.RegisterResponder(requestType, "https://"+configuration.Host+endpoint,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(statusCode, expectedResponse)

			resp.Header.Add("Content-Type", givenRequestHeaders["Content-Type"])
			resp.Header.Add("server", givenRequestHeaders["server"])
			resp.Header.Add("x-request-id", givenRequestHeaders["x-request-id"])
			resp.Header.Add("User-Agent", configuration.UserAgent)

			return resp, nil
		},
	)
}

func ValidateExpectedRequestBodiesMatches(t *testing.T, expectedRequest string, actualRequest string) {
	require.JSONEq(t, expectedRequest, actualRequest)
}

func ValidateExpectedRequestHeadersMatches(t *testing.T, headers http.Header) {
	assert.Equal(t, givenRequestHeaders["server"], headers.Get("server"))
	assert.Equal(t, givenRequestHeaders["Content-Type"], headers.Get("Content-Type"))
	assert.Equal(t, givenRequestHeaders["x-request-id"], headers.Get("x-request-id"))
	assert.Equal(t, configuration.UserAgent, headers.Get("User-Agent"))
}
