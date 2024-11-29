package cmd

import (
	"fmt"

	"github.com/patrickcping/davinci-pingcli/internal/utils"
	"github.com/spf13/cobra"
)

const (
	connectorsCmdName = "connectors"
)

var connectorsCmd = &cobra.Command{
	Use:     connectorsCmdName,
	Aliases: []string{utils.RemoveTrailingS(connectorsCmdName)},
	Short:   "Operations on DaVinci connector definitions",
	Long: fmt.Sprintf(`Provides command operations on DaVinci connectors.
	`),
}

func init() {
	// General function commands
	connectorsCmd.AddCommand(
		connectorsSchemaCmd,
	)
}
