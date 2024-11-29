package cmd

import (
	"fmt"

	"github.com/patrickcping/davinci-pingcli/internal/utils"
	"github.com/spf13/cobra"
)

const (
	flowsCmdName = "flows"
)

var flowsCmd = &cobra.Command{
	Use:     flowsCmdName,
	Aliases: []string{utils.RemoveTrailingS(flowsCmdName)},
	Short:   "Operations on DaVinci flows",
	Long: fmt.Sprintf(`Provides command operations on DaVinci flows.

	Examples:
	
	List flows using path parameters:
		davinci-pingcli %[1]s %[2]s --username myuser --password mypassword --admin-environment-id 00...00 --environment-id 00...00 --region Europe

	List flows using environment variables:
		export PINGCLI_DAVINCI_USERNAME=myuser
		export PINGCLI_DAVINCI_PASSWORD=mypassword
		export PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_REGION=Europe
		davinci-pingcli %[1]s %[2]s

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
	// General function commands
	flowsCmd.AddCommand(
		flowsListCmd,
		flowsVersionsCmd,
	)
}
