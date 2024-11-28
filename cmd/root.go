package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/patrickcping/davinci-pingcli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	adminUsernameParamLong  = "username"
	adminUsernameParamShort = "u"

	adminPasswordParamLong  = "password"
	adminPasswordParamShort = "P"

	adminEnvironmentIdParamLong  = "admin-environment-id"
	adminEnvironmentIdParamShort = "e"

	environmentIdParamLong  = "environment-id"
	environmentIdParamShort = "t"

	regionParamLong  = "region"
	regionParamShort = "r"
)

var (
	davinciConfigKeys = []string{
		adminUsernameParamLong,
		adminPasswordParamLong,
		adminEnvironmentIdParamLong,
		environmentIdParamLong,
		regionParamLong,
	}
)

var (
	adminUsername      string
	adminPassword      string
	adminEnvironmentId string
	environmentId      string
	region             string
)

const (
	configEnvPrefix = "PINGCLI_DAVINCI"
)

var (
	configKeys = map[string]string{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "davinci-pingcli",
	Short: "davinci-pingcli is a CLI utility to assist with using PingOne DaVinci through command line.",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		l := logger.Get()

		err := initConfig()
		if err != nil {
			return err
		}

		err = bindParams(cmd)
		if err != nil {
			return err
		}

		l.Debug().Msgf("PreRun Command called.")

		slog.Debug("PreRun command done")

		return nil
	},
}

func SetVersionInfo(version, commit string) {
	rootCmd.Version = fmt.Sprintf("%s (Git Commit SHA %s)", version, commit)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	l := logger.Get()

	// General function commands
	rootCmd.AddCommand(
		flowsCmd,
	)

	for _, v := range davinciConfigKeys {
		configKeyVal := strings.ReplaceAll(v, "-", "")
		configKeys[configKeyVal] = fmt.Sprintf("%s", configKeyVal)
	}

	// Add config flags
	rootCmd.PersistentFlags().StringVarP(&adminUsername, adminUsernameParamLong, adminUsernameParamShort, "", "The admin username used to connect to DaVinci.")
	if err := rootCmd.MarkPersistentFlagRequired(adminUsernameParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", adminUsernameParamLong)
	}
	rootCmd.PersistentFlags().StringVarP(&adminPassword, adminPasswordParamLong, adminPasswordParamShort, "", "The admin password used to connect to DaVinci.")
	if err := rootCmd.MarkPersistentFlagRequired(adminPasswordParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", adminPasswordParamLong)
	}
	rootCmd.PersistentFlags().StringVarP(&adminEnvironmentId, adminEnvironmentIdParamLong, adminEnvironmentIdParamShort, "", "The PingOne environment ID that contains the admin user.")
	if err := rootCmd.MarkPersistentFlagRequired(adminEnvironmentIdParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", adminEnvironmentIdParamLong)
	}
	rootCmd.PersistentFlags().StringVarP(&environmentId, environmentIdParamLong, environmentIdParamShort, "", "The PingOne environment ID to control configuration for.")
	if err := rootCmd.MarkPersistentFlagRequired(environmentIdParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", environmentIdParamLong)
	}
	rootCmd.PersistentFlags().StringVarP(&region, regionParamLong, regionParamShort, "", "The region where the PingOne environment is located.  Options are AsiaPacific, Canada, Europe and NorthAmerica.")
	if err := rootCmd.MarkPersistentFlagRequired(regionParamLong); err != nil {
		l.Err(err).Msgf("Error marking flag %s as required.", regionParamLong)
	}
}

func initConfig() error {
	l := logger.Get()

	l.Debug().Msgf("Initialising configuration..")

	viper.SetConfigName(".davinci-pingcli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	l.Debug().Msgf("Reading configuration..")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			l.Err(err).Msgf("Error reading configuration file.")
			return err
		}
	}

	viper.SetEnvPrefix(configEnvPrefix)

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	viper.AutomaticEnv()

	l.Debug().Msgf("Setting configuration..")

	return nil
}

func bindParams(cmd *cobra.Command) error {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		configName := strings.ReplaceAll(f.Name, "-", "")

		if v, ok := configKeys[configName]; ok {
			configName = v
		}

		if !f.Changed && viper.IsSet(configName) {
			// if err = cmd.Flags().SetAnnotation(f.Name, cobra.BashCompOneRequiredFlag, []string{"false"}); err != nil {
			// 	l.Err(err).Msgf("Error setting required status for flag %s", f.Name)
			// 	return
			// }
			viperValue := viper.Get(configName)
			switch v := viperValue.(type) {
			case []interface{}:
				values := make([]string, 0)
				for _, val := range v {
					values = append(values, fmt.Sprintf("%v", val))
				}
				cmd.Flags().Set(f.Name, strings.Join(values, ","))
			default:
				cmd.Flags().Set(f.Name, fmt.Sprintf("%v", v))
			}
		}
	})

	return nil
}
