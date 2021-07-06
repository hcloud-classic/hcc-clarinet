package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"hcc/clarinet/lib/config"
)

// aipCmd represents the aip command
var logoutCmd = &cobra.Command{
	Use:     "logout",
	Short:   "User logout",
	Long:    ``,
	Args:    cobra.MinimumNArgs(0),
	PreRunE: checkToken,
	Run: func(cmd *cobra.Command, args []string) {
		token := config.User.Token
		if token != "" {
			config.RemoveTokenString()
		} else {
			fmt.Println("Please log in first")
		}

	},
}
