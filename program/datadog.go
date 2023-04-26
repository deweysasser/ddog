package program

import (
	"context"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/rs/zerolog/log"
)

func (program *Options) Client(ctx context.Context) (*datadog.APIClient, context.Context) {

	//ctx = datadog.NewDefaultContext(ctx)

	ctx = context.WithValue(
		ctx,
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: program.DatadogApiKey,
			},
			"appKeyAuth": {
				Key: program.DatadogAppKey,
			},
		},
	)

	log.Debug().
		Msg("Creating Datadog client")
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	return apiClient, ctx
}
