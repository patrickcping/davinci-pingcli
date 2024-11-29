package davinci

import (
	"context"
	"log/slog"

	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type DaVinciEnvironment struct {
	EnvironmentID string
	Client        *davinci.APIClient
}

func (r DaVinciEnvironment) DeleteFlowVersion(ctx context.Context, flowID, flowVersionID string) error {
	// Run the API call
	flowsResult, err := r.Client.DeleteFlowVersion(r.EnvironmentID, flowID, flowVersionID)

	if err != nil {
		return err
	}

	slog.Debug("Delete flow version response", "response", flowsResult, "flowID", flowID, "flowVersionID", flowVersionID)

	return nil
}
