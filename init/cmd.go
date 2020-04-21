package init

import (
	"hcc/clarinet/cmd"
	"os"

	"github.com/spf13/cobra"
)

func cmdInit() error {
	cmd.ReadyServerCmd()

	var rootCmd = &cobra.Command{Use: os.Args[0]}
	rootCmd.AddCommand(cmd.ServerCmd)
	return rootCmd.Execute()
}
