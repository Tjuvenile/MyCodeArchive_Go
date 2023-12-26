package cmd

import (
	"MyCodeArchive_Go/utils/tool"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "dcs",
	Short: "smart connect operations for unstructured storage service",
	Long: `smart connect enable client connection load balancing and dynamic failover 
and failback of client connections across storage gateway nodes to optimize 
use of cluster resources`,
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		tool.PrintFailedMsg(err.Error(), false)
		fmt.Println()
		cmd, _, _ := RootCmd.Find(os.Args[1:])
		fmt.Printf(cmd.UsageString())
	}
}
