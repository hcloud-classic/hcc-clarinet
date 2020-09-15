package init

import (
	"github.com/spf13/cobra"
	"hcc/clarinet/cmd"
)

func cmdInit() error {
	cmd.ReadyServerCmd()

	var rootCmd = &cobra.Command{Use: "clarinet"}
	rootCmd.AddCommand(cmd.ServerCmd)
	return rootCmd.Execute()
}
