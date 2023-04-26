package program

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
)

type Monitor struct {
	SaveTo string `group:"save" help:"Save the current monitor configuration to a file" type:"existingdir"`
}

// Run runs the program
func (program *Monitor) Run(options *Options) error {
	// Configure Datadog client

	ctx := context.Background()

	client, ctx := options.Client(ctx)

	// Save monitors to JSON file

	api := datadogV1.NewMonitorsApi(client)

	log.Debug().Msg("Listing api")
	monitors, resp, err := api.ListMonitors(ctx, *datadogV1.NewListMonitorsOptionalParameters())

	if err != nil {
		log.Debug().
			Int("status_code", resp.StatusCode).
			Str("status", resp.Status).
			Str("context", fmt.Sprint(ctx)).
			Msg("Response")
		return errors.Wrap(err, "Error listing api")

	}

	for _, monitor := range monitors {

		name := monitor.GetName()
		id := monitor.GetId()

		bytes, err := json.MarshalIndent(monitor, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshalling monitor %s (%d): %w", name, id, err)
		}

		os.WriteFile(fmt.Sprintf("%s/%s.json", program.SaveTo, name), bytes, 0644)
	}

	return nil
}
