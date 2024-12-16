package cli

import (
	"fmt"

	"github.com/mytechpal-net/dmarc-analyzer/webservice/pkg/cli/dmarc"
	"github.com/spf13/cobra"
)

var dmarcCmd = &cobra.Command{
	Use:   "dmarc",
	Short: "dmarc command",
	Run: func(cmd *cobra.Command, args []string) {
		dmarcCommandExcutor()
	},
}

func dmarcCommandExcutor() {
	fmt.Println("Running dmarc command...")
}

func init() {
	dmarcCmd.AddCommand(dmarc.ParserCmd)
	dmarcCmd.AddCommand(dmarc.DbCmd)
}
