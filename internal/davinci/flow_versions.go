package davinci

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
)

func (r DaVinciEnvironment) DeleteFlowVersion(ctx context.Context, flowID, flowVersionID string) error {

	if flowID == "" {
		return fmt.Errorf("Flow ID is required")
	}

	if flowVersionID == "" {
		return fmt.Errorf("Flow Version ID is required")
	}

	// Get the flow details
	flowResult, err := r.Client.ReadFlow(r.EnvironmentID, flowID)
	if err != nil {
		return err
	}

	// If the version is the current or deployed flow, then fail
	versionIDInt, err := strconv.ParseInt(flowVersionID, 10, 32)
	if err != nil {
		return err
	}

	if v := flowResult.Flow.CurrentVersion; v != nil && int32(versionIDInt) == *v {
		return fmt.Errorf("Cannot delete the current version of a flow")
	}

	if v := flowResult.Flow.PublishedVersion; v != nil && int32(versionIDInt) == *v {
		return fmt.Errorf("Cannot delete the deployed version of a flow")
	}

	// Run the API call
	deleteFlowVersionResult, err := r.Client.DeleteFlowVersion(r.EnvironmentID, flowID, flowVersionID)
	if err != nil {
		return err
	}

	slog.Debug("Delete flow version response", "response", deleteFlowVersionResult, "flowID", flowID, "flowVersionID", flowVersionID)

	return nil
}
