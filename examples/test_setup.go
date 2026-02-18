package examples

import (
	"context"
	"sync"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
)

const (
	baseURL = "<BASE_URL>"
	apiKey  = "<API_KEY>"
)

var (
	once         sync.Once
	sharedClient *api.APIClient
	sharedAuth   context.Context
)

func newClientAndAuth() (*api.APIClient, context.Context) {
	once.Do(func() {
		configuration := infobip.NewConfiguration()
		configuration.Host = baseURL

		sharedClient = api.NewAPIClient(configuration)
		sharedAuth = context.WithValue(
			context.Background(),
			infobip.ContextAPIKeys,
			map[string]infobip.APIKey{
				"APIKeyHeader": {Key: apiKey},
			},
		)
	})

	return sharedClient, sharedAuth
}
