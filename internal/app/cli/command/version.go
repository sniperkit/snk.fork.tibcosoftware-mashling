/*
Sniperkit-Bot
- Status: analyzed
*/

package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sniperkit/snk.fork.tibcosoftware-mashling/internal/app/version"
)

func init() {
	cliCommand.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the mashling-cli version",
	Long:  `Prints the mashling-cli version and build details`,
	Run:   ver,
}

func ver(command *cobra.Command, args []string) {
	fmt.Println("Version: ", version.Version)
	fmt.Println("Build Date: ", version.BuildDate)
}
