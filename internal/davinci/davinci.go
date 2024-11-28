package davinci

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/fatih/color"
	"github.com/patrickcping/davinci-pingcli/internal/output"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type DaVinciEnvironment struct {
	EnvironmentID string
	Client        *davinci.APIClient
}

func (r DaVinciEnvironment) ReadFlows(ctx context.Context) error {
	// Run the API call
	flowsResult, err := r.Client.ReadFlows(r.EnvironmentID, &davinci.Params{
		Limit: "200",
	})

	if err != nil {
		return err
	}

	// output here
	listStr := "Flows:\n"

	// We need to enable/disable colorize before applying the color to the string below.
	nameFmt := color.New(color.Bold).SprintFunc()
	deployedFmt := color.New(color.Bold, color.FgGreen).SprintFunc()
	neverDeployedFmt := color.New(color.Bold, color.FgRed).SprintFunc()

	for _, flow := range flowsResult {
		if flow.DeployedDate == nil {
			listStr += "- " + nameFmt(flow.Name) + neverDeployedFmt(" (never deployed)") + " \n"
		} else {
			listStr += "- " + nameFmt(flow.Name) + deployedFmt(" (deployed)") + "\n"
		}

		if v := flow.CurrentVersion; v != nil {
			listStr += fmt.Sprintf("\tCurrent Version:\t%d\n", *v)
		}

		listStr += "\tDeployed Version:\t"
		if v := flow.PublishedVersion; v != nil {
			listStr += fmt.Sprintf("%d", *v)

			if d := flow.DeployedDate; d != nil {
				listStr += fmt.Sprintf(" (%s)", d)
			}
		} else {
			listStr += neverDeployedFmt("Never deployed")

		}

		listStr += "\n"
		if v := flow.Description; v != nil && *v != "" {
			listStr += fmt.Sprintf("\tDescription:\n\t\t%s\n", strings.TrimSpace(*v))
		}

		listStr += "\n"
	}

	output.Message(listStr, nil)

	return nil
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
