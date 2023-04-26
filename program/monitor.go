package program

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"regexp"
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

	return program.SaveMonitors(monitors)
}

var invalidFilenameCharacters = regexp.MustCompile(`[^a-zA-Z0-9\-_]+`)
var filenameLength = 200

// SaveMonitors saves the given monitors to the directory in program.SaveTo
func (program *Monitor) SaveMonitors(monitors []datadogV1.Monitor) error {
	for _, monitor := range monitors {
		name := monitor.GetName()
		id := monitor.GetId()

		//Extract all text between {{#is_alert}} and {{/is_alert}}
		name = extractAlertText(name)

		token := invalidFilenameCharacters.ReplaceAllString(name, "_")
		if token[0] == '_' {
			token = token[1:]
		}

		if len(token) > filenameLength {
			token = token[0:filenameLength]
		}

		bytes, err := json.MarshalIndent(monitor, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshalling monitor %s (%d): %w", token, id, err)
		}

		outputFile := fmt.Sprintf("%s/%s-%d.json", program.SaveTo, token, id)

		log.Debug().
			Str("name", name).
			Int64("id", id).
			Str("output_file", outputFile).
			Msg("Saving monitor")

		if err = os.WriteFile(outputFile, bytes, 0644); err != nil {
			return fmt.Errorf("error saving monitor %s (%d): %w", name, id, err)
		}
	}

	return nil
}

// extractAlertText extracts all text between {{#is_alert}} and {{/is_alert}}
func extractAlertText(name string) string {

	re := regexp.MustCompile(`{{#is_alert}}(.*?){{/is_alert}}`)
	match := re.FindStringSubmatch(name)
	if len(match) > 0 {
		return match[1]
	}
	return name
}
