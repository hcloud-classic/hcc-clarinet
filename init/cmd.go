package init

import (
	"github.com/spf13/cobra"
	"hcc/clarinet/cmd"
)

func cmdInit() error {
	cmd.ReadyServerCmd()
	cmd.ReadyNodeCmd()
	cmd.ReadySubnetCmd()

	var rootCmd = &cobra.Command{Use: "clarinet"}
	rootCmd.AddCommand(cmd.ServerCmd)
	rootCmd.AddCommand(cmd.NodeCmd)
	rootCmd.AddCommand(cmd.SubnetCmd)
	return rootCmd.Execute()
}
