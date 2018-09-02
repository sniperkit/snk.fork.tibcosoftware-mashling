/*
Sniperkit-Bot
- Status: analyzed
*/

package command

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	gwerrors "github.com/sniperkit/snk.fork.tibcosoftware-mashling/internal/pkg/model/errors"
)

func init() {
	cliCommand.AddCommand(validateCommand)
}

var validateCommand = &cobra.Command{
	Use:   "validate",
	Short: "Validates a mashling.json configuration file",
	Long:  `Validates a provided mashling.json configuration file based off of the supported Mashling schema versions`,
	Run:   validate,
}

func validate(command *cobra.Command, args []string) {
	err := loadGateway()
	if err != nil {
		log.Println("Invalid configuration file!:", err)
		for _, errd := range gateway.Errors() {
			switch e := errd.(type) {
			case *gwerrors.UndefinedReference:
				log.Printf("%s: %s", e.Type(), e.Details())
			case *gwerrors.MissingDependency:
				log.Println("Missing dependencies found: ", strings.Join(e.MissingDependencies, " "))
			default:
				log.Printf("Do not know how to handle error type %T!\n", e)
			}
		}
		os.Exit(1)
	} else {
		log.Println("Valid.")
	}
}
