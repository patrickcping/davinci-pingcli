package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/patrickcping/davinci-pingcli/internal/davinci"
	"github.com/patrickcping/davinci-pingcli/internal/logger"
	dvsdk "github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/spf13/cobra"
)

const (
	connectorsSchemaCmdName = "schema"
)

const (
	outputFormatJSONParamLong = "json"
)

var (
	outputFormatJSON bool
)

var connectorsSchemaCmd = &cobra.Command{
	Use:   connectorsSchemaCmdName,
	Short: "Prints out a JSON representation of the current DaVinci connector definitions in an environment.",
	Long: fmt.Sprintf(`Prints out a JSON representation of the current DaVinci connector definitions in an environment.

	Examples:
	
	Using path parameters:
		davinci-pingcli %[1]s %[2]s --username myuser --password mypassword --admin-environment-id 00...00 --environment-id 00...00 --region Europe --json

	Using environment variables:
		export PINGCLI_DAVINCI_USERNAME=myuser
		export PINGCLI_DAVINCI_PASSWORD=mypassword
		export PINGCLI_DAVINCI_ADMIN_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_ENVIRONMENT_ID=00...00
		export PINGCLI_DAVINCI_REGION=Europe
		davinci-pingcli %[1]s %[2]s --json
`, connectorsCmdName, connectorsSchemaCmdName),
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.Get()
		l.Debug().Msgf("connectors schema Command called.")

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
			log.Fatalf("Error creating the DaVinci client: %s", err)
		}

		dvEnvironment := davinci.DaVinciEnvironment{
			EnvironmentID: environmentId,
			Client:        apiClient,
		}

		err = dvEnvironment.ConnectorSchema(cmd.Root().Context(), outputFormatJSON)
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	},
}

func init() {
	l := logger.Get()

	connectorsSchemaCmd.Flags().BoolVar(&outputFormatJSON, outputFormatJSONParamLong, true, "Output the result in JSON format.")
	if err := rootCmd.MarkFlagRequired(outputFormatJSONParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", outputFormatJSONParamLong)
	}

}
