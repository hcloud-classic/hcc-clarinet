package cmd

import (
	"github.com/spf13/cobra"
)

var Cmd *cobra.Command = nil

func Init() {
	ReadyServerCmd()
	ReadyNodeCmd()
	ReadySubnetCmd()
	ReadyAIPCmd()
	ReadyConfigCmd()

	Cmd = &cobra.Command{Use: "clarinet"}
	Cmd.AddCommand(serverCmd, nodeCmd, subnetCmd, aipCmd, configCmd)
}
