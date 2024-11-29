package cmd

import (
	"fmt"
	"os"

	"github.com/patrickcping/davinci-pingcli/internal/davinci"
	"github.com/patrickcping/davinci-pingcli/internal/logger"
	"github.com/patrickcping/davinci-pingcli/internal/output"
	dvsdk "github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/spf13/cobra"
)

const (
	flowsListCmdName = "list"
)

var flowsListCmd = &cobra.Command{
	Use:   flowsListCmdName,
	Short: "Lists the DaVinci flows in an environment.",
	Long: fmt.Sprintf(`Lists the DaVinci flows in an environment, along with the latest version and latest deployed version.

	Examples:
	
	Using path parameters:
		davinci-pingcli %[1]s %[2]s --username myuser --password mypassword --admin-environment-id 00...00 --environment-id 00...00 --region Europe

	Using environment variables:
		export PINGCLI_DAVINCI_USERNAME=myuser
		export PINGCLI_DAVINCI_PASSWORD=mypassword
		export PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_REGION=Europe
		davinci-pingcli %[1]s %[2]s
`, flowsCmdName, flowsListCmdName),
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

		err = dvEnvironment.ReadFlows(cmd.Root().Context())
		if err != nil {
			output.UserFatal(fmt.Sprintf("%s", err), map[string]interface{}{})
		}

		os.Exit(0)
	},
}
