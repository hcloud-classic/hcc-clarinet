package init

import (
	"github.com/spf13/cobra"
	"hcc/clarinet/cmd"
	"os"
)

func cmdInit() error {
	cmd.ReadyServerCmd()

	var rootCmd = &cobra.Command{Use: os.Args[0]}
	rootCmd.AddCommand(cmd.ServerCmd)
	return rootCmd.Execute()
}