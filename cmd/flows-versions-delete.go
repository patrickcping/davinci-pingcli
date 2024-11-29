package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/patrickcping/davinci-pingcli/internal/davinci"
	"github.com/patrickcping/davinci-pingcli/internal/logger"
	"github.com/patrickcping/davinci-pingcli/internal/output"
	dvsdk "github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/spf13/cobra"
)

const (
	flowsVersionsDeleteCmdName = "delete"
)

var (
	flowID        string
	flowVersionID string
	confirmDelete bool
)

var flowsVersionsDeleteCmd = &cobra.Command{
	Use:   fmt.Sprintf("%s --flow-id 00...00 --flow-version-id 10", flowsVersionsDeleteCmdName),
	Short: "Lists the DaVinci flows in an environment.",
	Long: fmt.Sprintf(`Lists the DaVinci flows in an environment, along with the latest version and latest deployed version.

	Using path parameters:
		davinci-pingcli %[1]s %[2]s %[3]s --flow-id 00...00 --flow-version-id 10 --username myuser --password mypassword --admin-environment-id 00...00 --environment-id 00...00 --region Europe

	Using environment variables:
		export PINGCLI_DAVINCI_USERNAME=myuser
		export PINGCLI_DAVINCI_PASSWORD=mypassword
		export PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_REGION=Europe
		davinci-pingcli %[1]s %[2]s %[3]s --flow-id 00...00 --flow-version-id 10 
	`, flowsCmdName, flowsVersionsCmdName, flowsVersionsDeleteCmdName),
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.Get()
		l.Debug().Msgf("flows list Command called.")

		//

		userAgent := fmt.Sprintf("patrickcping/davinci-pingcli/%s", cmd.Root().Version)

		cInput := dvsdk.ClientInput{
			Username:        adminUsername,
			Password:        adminPassword,
			PingOneRegion:   region,
			PingOneSSOEnvId: adminEnvironmentId,
			UserAgent:       userAgent,
		}
		apiClient, err := dvsdk.NewClient(&cInput)
		if err != nil {
			output.UserFatal("Error creating the DaVinci client", map[string]interface{}{
				"error": err,
			})
		}

		dvEnvironment := davinci.DaVinciEnvironment{
			EnvironmentID: environmentId,
			Client:        apiClient,
		}

		if flowID == "" || flowVersionID == "" {
			output.UserFatal("Flow ID and Flow Version ID are required.", map[string]interface{}{
				"flow ID":         flowID,
				"flow version ID": flowVersionID,
			})
		}

		if !confirmDelete {
			// Prompt for confirmation
			p := promptui.Prompt{
				Label:     fmt.Sprintf("Are you sure you want to delete the flow version %s for flow with ID %s?", flowVersionID, flowID),
				IsConfirm: true,
				Stdin:     os.Stdin,
			}

			// This is odd behavior discussed in https://github.com/manifoldco/promptui/issues/81
			// If err is type promptui.ErrAbort, the user can be assumed to have responded "No"
			confirmResult, err := p.Run()
			if err != nil {
				if errors.Is(err, promptui.ErrAbort) {
					confirmDelete = false
				}
				confirmDelete = false
			}

			if strings.ToUpper(confirmResult) == "Y" {
				confirmDelete = true
			}
		}

		if confirmDelete {
			output.Message("Deletion confirmed, deleting", map[string]interface{}{})
			err = dvEnvironment.DeleteFlowVersion(cmd.Root().Context(), flowID, flowVersionID)
			if err != nil {
				output.UserFatal(fmt.Sprintf("%s", err), map[string]interface{}{})
			}
			output.Message(fmt.Sprintf("Deletion of version %s of flow %s completed", flowVersionID, flowID), map[string]interface{}{})
		} else {
			output.Message("Deletion not confirmed. Exiting.", map[string]interface{}{})
		}

		os.Exit(0)
	},
}

func init() {
	l := logger.Get()

	// Add config flags
	flowsVersionsDeleteCmd.Flags().StringVar(&flowVersionID, "flow-version-id", "", "The flow version ID to delete a version from.")
	if err := flowsVersionsDeleteCmd.MarkFlagRequired("flow-version-id"); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", "flow-version-id")
	}

	// Add config flags
	flowsVersionsDeleteCmd.Flags().BoolVarP(&confirmDelete, "yes", "y", false, "Auto-confirm the deletion.")
}
