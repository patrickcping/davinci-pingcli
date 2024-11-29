package davinci

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/patrickcping/davinci-pingcli/internal/output"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func (r DaVinciEnvironment) ConnectorSchema(ctx context.Context, outputFormatJSON bool) error {
	// Run the API call
	connectors, err := r.Client.ReadConnectors(&r.EnvironmentID, nil)

	if err != nil {
		return err
	}

	bytes, err := davinci.Marshal(connectors, davinci.ExportCmpOpts{
		IgnoreEnvironmentMetadata: true,
	})
	if err != nil {
		return err
	}

	var sanitisedConnectors []davinci.Connector

	err = davinci.Unmarshal(bytes, &sanitisedConnectors, davinci.ExportCmpOpts{
		IgnoreEnvironmentMetadata: true,
	})
	if err != nil {
		return err
	}

	if outputFormatJSON {

		bytes, err := json.MarshalIndent(sanitisedConnectors, "", "    ")
		if err != nil {
			return err
		}

		output.Message(string(bytes[:]), nil)
	} else {
		return fmt.Errorf("Output format not supported")
	}

	return nil
}
