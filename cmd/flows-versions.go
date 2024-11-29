package cmd

import (
	"fmt"

	"github.com/patrickcping/davinci-pingcli/internal/logger"
	"github.com/patrickcping/davinci-pingcli/internal/utils"
	"github.com/spf13/cobra"
)

const (
	flowsVersionsCmdName = "versions"
)

var flowsVersionsCmd = &cobra.Command{
	Use:     flowsVersionsCmdName,
	Aliases: []string{utils.RemoveTrailingS(flowsVersionsCmdName)},
	Short:   "Operations on DaVinci flow versions",
	Long: fmt.Sprintf(`Provides command operations on DaVinci flow versions.

	Examples:
	
	Delete flow version using path parameters:
		davinci-pingcli %[1]s %[2]s %[3]s --flow-id 00...00 --flow-version-id 10 --username myuser --password mypassword --admin-environment-id 00...00 --environment-id 00...00 --region Europe

	Delete flow version using environment variables:
		export PINGCLI_DAVINCI_USERNAME=myuser
		export PINGCLI_DAVINCI_PASSWORD=mypassword
		export PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_REGION=Europe
		davinci-pingcli %[1]s %[2]s %[3]s --flow-id 00...00 --flow-version-id 10
	`, flowsCmdName, flowsVersionsCmdName, flowsVersionsDeleteCmdName),
}

func init() {
	l := logger.Get()

	// General function commands
	flowsVersionsCmd.AddCommand(
		flowsVersionsDeleteCmd,
	)

	flowsVersionsCmd.PersistentFlags().StringVarP(&flowID, "flow-id", "", "", "The flow ID to manage a version for.")
	if err := flowsVersionsCmd.MarkPersistentFlagRequired("flow-id"); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", "flow-id")
	}
}
