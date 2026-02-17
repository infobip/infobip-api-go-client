package infobip

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
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

func SetUpGetRequestWithFileContent(endpoint string, expectedResponseContent []byte, statusCode int) {
	httpmock.RegisterResponder("GET", "https://"+configuration.Host+endpoint,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(statusCode, expectedResponseContent)

			resp.Header.Add("Content-Type", "audio/vnd.wave")
			resp.Header.Add("server", givenRequestHeaders["server"])
			resp.Header.Add("x-request-id", givenRequestHeaders["x-request-id"])
			resp.Header.Add("User-Agent", configuration.UserAgent)

			return resp, nil
		},
	)
}

func SetUpGetRequest(url string, givenResponse []byte, statusCode int) {
	httpmock.RegisterResponder("GET", "https://"+configuration.Host+url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(statusCode, givenResponse)
			resp.Header.Set("Content-Type", "audio/vnd.wave")
			resp.Header.Add("Server", givenRequestHeaders["server"])
			resp.Header.Add("X-Request-Id", givenRequestHeaders["x-request-id"])
			resp.Header.Add("User-Agent", configuration.UserAgent)

			return resp, nil
		},
	)
}

func SetUpGetRequestOctet(url string, givenResponse []byte, statusCode int) {
	httpmock.RegisterResponder("GET", "https://"+configuration.Host+url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewBytesResponse(statusCode, givenResponse)
			resp.Header.Set("Content-Type", "application/octet-stream")
			resp.Header.Add("Server", givenRequestHeaders["server"])
			resp.Header.Add("X-Request-Id", givenRequestHeaders["x-request-id"])
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

func TestConfigurationServerURLWithContext(t *testing.T) {
	tests := []struct {
		name    string
		host    string
		want    string
		wantErr string
	}{
		{
			name: "prepends https when missing",
			host: "api.infobip.com",
			want: "https://api.infobip.com",
		},
		{
			name: "keeps https scheme",
			host: "https://api.infobip.com",
			want: "https://api.infobip.com",
		},
		{
			name: "trims trailing slashes",
			host: "https://api.infobip.com/",
			want: "https://api.infobip.com",
		},
		{
			name: "allows localhost without scheme (defaults to https)",
			host: "localhost:8080",
			want: "https://localhost:8080",
		},
		{
			name: "allows localhost with http scheme",
			host: "http://localhost:8080",
			want: "http://localhost:8080",
		},
		{
			name: "allows 127.0.0.1 with scheme",
			host: "http://127.0.0.1:8080",
			want: "http://127.0.0.1:8080",
		},
		{
			name: "trims spaces",
			host: "  api.infobip.com  ",
			want: "https://api.infobip.com",
		},
		{
			name:    "rejects http for non-localhost",
			host:    "http://api.infobip.com",
			wantErr: "Only HTTPS is supported for the base URL",
		},
		{
			name:    "rejects empty",
			host:    "   ",
			wantErr: "Base URL must not be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := infobip.NewConfiguration()
			cfg.Host = tt.host

			got, err := cfg.ServerURLWithContext(context.Background(), "")
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}
				if err.Error() != tt.wantErr {
					t.Fatalf("expected error %q, got %q", tt.wantErr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestAPIKeyAuthAcceptsRawString(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = Host

	client := api.NewAPIClient(configuration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	endpoint := "/callrouting/1/routes"
	expected := `{"results":[],"paging":{"page":0,"size":0,"totalPages":0,"totalResults":0}}`
	rawKey := strings.TrimPrefix(ApiKey, "App ")

	httpmock.RegisterResponder("GET", "https://"+configuration.Host+endpoint,
		func(req *http.Request) (*http.Response, error) {
			if got := req.Header.Get("Authorization"); got != "App "+rawKey {
				return httpmock.NewStringResponse(400, ""), fmt.Errorf("unexpected Authorization header: %s", got)
			}
			response := httpmock.NewStringResponse(200, expected)
			response.Header.Set("Content-Type", "application/json")
			return response, nil
		},
	)

	ctx := context.WithValue(context.Background(), infobip.ContextAPIKeys, rawKey)

	resp, _, err := client.
		CallRoutingAPI.
		GetCallRoutes(ctx).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatalf("expected response, got nil")
	}
}
