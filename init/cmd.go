package init

import (
	"github.com/spf13/cobra"
	"hcc/clarinet/cmd"
)

func cmdInit() error {
	cmd.ReadyServerCmd()
	cmd.ReadyNodeCmd()
	cmd.ReadySubnetCmd()
	cmd.ReadyAIPCmd()

	var rootCmd = &cobra.Command{Use: "clarinet"}
	rootCmd.AddCommand(cmd.ServerCmd, cmd.NodeCmd, cmd.SubnetCmd, cmd.AIPCmd)
	return rootCmd.Execute()
}
